package main

import (
	"bufio"
	"bytes"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"io"
	"log"
	"net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/5/4 22:29
 * @Desc:	监控本机mac lo0网卡流量
 * 				- 外部HTTP服务（经过网卡可以监听到）：goinfra/http/httpservertest/sample/main.go
 * 					- 需先启动外部http服务
 */

func MonitorMACWIFINoserverNoclient() {
	// 打开网络接口
	// lo0网卡
	// handle, err := pcap.OpenLive("lo0", 65535, true, pcap.BlockForever)
	// eno网卡：wifi网卡
	handle, err := pcap.OpenLive("en0", 65535, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// 设置过滤条件
	filter := "host ug.baidu.com" // 注意切换网卡。需使用lo0网卡`handle, err := pcap.OpenLive("lo0", 65535, true, pcap.BlockForever)`
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}

	// 开始监听网络流量
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

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
		log.Printf("tcp.SrcPort:%v;tcp.DstPort:%v", tcp.SrcPort, tcp.DstPort)
		log.Printf("tcp.SYN:%v;tcp.ACK:%v;tcp.PSH:%v;tcp.FIN:%v;", tcp.SYN, tcp.ACK, tcp.PSH, tcp.FIN)
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
		listenPort := layers.TCPPort(443)
		log.Println("监听的端口layers.TCPPort:", listenPort)
		if tcp.SrcPort == listenPort || tcp.DstPort == listenPort {
			log.Println("tcp.SrcPort == listenPort || tcp.DstPort == listenPort")
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
				log.Println("Response body Content-Length:", len(body))
			} else {
				log.Println("Response err != nil:", err)
			}
		}
	}
}
