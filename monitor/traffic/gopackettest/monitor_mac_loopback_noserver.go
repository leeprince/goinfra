package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/5/4 22:29
 * @Desc:	监控本机mac lo0网卡流量
 * 				- 该应用服务自行模拟客户端,外部HTTP服务：goinfra/http/httpservertest/sample/main.go
 * 					- 需先启动外部http服务
 */

func MonitorMACLoopbackNoserver() {
	// 打开网络接口
	// lo0网卡
	handle, err := pcap.OpenLive("lo0", 65535, true, pcap.BlockForever)
	// eno网卡：wifi网卡
	// handle, err := pcap.OpenLive("en0", 65535, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// 设置过滤条件
	filter := "host localhost and port 8090" // 注意切换网卡。需使用lo0网卡`handle, err := pcap.OpenLive("lo0", 65535, true, pcap.BlockForever)`
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}

	// 开始监听网络流量
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	// --- 模拟HTTP客户端
	go func() {
		time.Sleep(time.Second * 1)
		log.Println("创建HTTP客户端...")

		// 创建HTTP客户端
		req, err := http.NewRequest("GET", "http://localhost:8090/prince/get", nil)
		if err != nil {
			log.Println("NewRequest err:", err)
			return
		}
		// 发送请求
		httpClient := http.DefaultClient
		resp, err := httpClient.Do(req)
		log.Println("httpClient.Do(req)", resp)
		if err != nil {
			log.Println("http.DefaultClient.Do", err)
			return
		}
		defer resp.Body.Close()
		bodyByte, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("http.DefaultClient.Do", err)
			return
		}
		log.Println("httpClient.Do(req)", string(bodyByte))
		log.Println("创建HTTP客户端...end")
	}()

	// --- 模拟HTTP客户端 -end

	log.Println("开始监听:", filter)
	for packet := range packetSource.Packets() {
		log.Println(">>>")
		log.Println(">>> >>>")
		log.Println(">>> >>> >>> range packetSource.Packets():", packet)

		// 解析数据包
		ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
		if ethernetLayer == nil {
			log.Println("packet.Layer(layers.LayerTypeEthernet) nil")
		}

		loopbackLayer := packet.Layer(layers.LayerTypeLoopback)
		if loopbackLayer == nil {
			log.Println("packet.Layer(layers.LayerTypeLoopback) nil")
		}

		ipV4Layer := packet.Layer(layers.LayerTypeIPv4)
		if ipV4Layer == nil {
			log.Println("packet.Layer(layers.LayerTypeIPv4) nil")
		} else {
			log.Println("layers.LayerTypeIPv4:", ipV4Layer)
			log.Println("layers.LayerTypeIPv4 LayerType:", ipV4Layer.LayerType())
			log.Println("layers.LayerTypeIPv4 LayerContents:", string(ipV4Layer.LayerContents()))
			log.Println("layers.LayerTypeIPv4 LayerPayload:", string(ipV4Layer.LayerPayload()))
			log.Println("<<<layers.LayerTypeIPv4:")
		}

		tcpLayer := packet.Layer(layers.LayerTypeTCP)
		if tcpLayer == nil {
			log.Println("packet.Layer(layers.LayerTypeTCP) nil")
		} else {
			log.Println("layers.LayerTypeTCP:", tcpLayer)
			log.Println("layers.LayerTypeTCP LayerType:", tcpLayer.LayerType())
			log.Println("layers.LayerTypeTCP LayerPayload:", string(tcpLayer.LayerPayload()))
			log.Println("<<<layers.LayerTypeTCP")
		}

		if ethernetLayer == nil && loopbackLayer == nil && ipV4Layer == nil && tcpLayer == nil {
			log.Println("ethernetLayer == nil && loopbackLayer == nil && ipV4Layer == nil && tcpLayer == nil")
			continue
		}

		tcp, ok := tcpLayer.(*layers.TCP)
		if !ok {
			log.Println("tcpLayer.(*layers.TCP) !ok:")
			continue
		}
		log.Println(">>>tcp:", tcp)
		log.Println("tcp.Payload:", string(tcp.Payload))
		log.Println("tcp.SrcPort:", tcp.SrcPort)
		log.Println("tcp.DstPort:", tcp.DstPort)
		log.Println("tcp.SYN:", tcp.SYN)
		log.Println("tcp.ACK:", tcp.ACK)
		log.Println("tcp.PSH:", tcp.PSH)
		log.Println("tcp.FIN:", tcp.FIN)
		log.Println("<<<tcp")
		tcpPayload := tcp.Payload

		// 解析应用层数据包;需要监听HTTP时，当有发生HTTP请求（已经在TCP握手后，发生的HTTP请求）时，appLayer也不可能为空，即tcpLayer不可能为空
		appLayer := packet.ApplicationLayer()
		if appLayer == nil {
			log.Println("packet.ApplicationLayer() appLayer nil")
			continue
		}
		log.Println("packet.ApplicationLayer():", appLayer)
		log.Println("packet.ApplicationLayer() LayerType:", appLayer.LayerType())
		log.Println("packet.ApplicationLayer() LayerContents:", string(appLayer.LayerContents()))
		log.Println("packet.ApplicationLayer() LayerPayload:", string(appLayer.LayerPayload()))
		log.Println("packet.ApplicationLayer() Payload:", string(appLayer.Payload()))
		log.Println("<<<packet.ApplicationLayer()")

		// 判断是否为HTTP请求或响应
		// 监听的端口
		listenPort := layers.TCPPort(8090)
		log.Println("监听的端口layers.TCPPort:", listenPort)
		if tcp.SrcPort == listenPort || tcp.DstPort == listenPort {
			// 尝试解析HTTP请求
			req, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(tcpPayload)))
			if err == nil {
				log.Println("ReadRequest Request req:", req)
				continue
			} else {
				log.Println("ReadRequest err != nil:", err)
			}

			// 尝试解析HTTP响应
			resp, err := http.ReadResponse(bufio.NewReader(bytes.NewReader(tcpPayload)), nil)
			if err == nil {
				log.Println("ReadResponse err == nil:", resp)

				body, err := io.ReadAll(resp.Body)
				if err != nil {
					log.Println("ReadResponse io.ReadAll(resp.Body) err != nil:", err)
					continue
				}
				log.Println("Response body:", string(body))
			} else {
				log.Println("Response err != nil:", err)
			}
		}

		if tcp.SYN && !tcp.ACK {
			log.Println("--- tcp.SYN && !tcp.ACK ---")
			requestPayload := "GET http://localhost:8091/prince/get HTTP/1.1"

			// 发 HTTP 请求
			if strings.Contains(string(tcp.Payload), requestPayload) {
				fmt.Println("--------------Sending request:", requestPayload)
			}
		} else if tcp.ACK && tcp.PSH && tcp.FIN {
			responsePayload := "HTTP/1.1 200 OK\r\n"
			log.Println("--- tcp.ACK && tcp.PSH && tcp.FIN ---")

			// 接收 HTTP 响应
			if strings.Contains(string(tcp.Payload), responsePayload) {
				fmt.Println("--------------Received response:", responsePayload)
				fmt.Println(string(tcp.Payload))
			}
		}
	}
}
