package wiresharkfile

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"net"
	"path/filepath"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/8 20:05
 * @Desc:
 */

func MonitorPcapngFileTcp(pcapngFilePath string) {
	// ---
	fileName := filepath.Base(pcapngFilePath)
	// ---
	
	// 打开 pcapng 文件
	handle, err := pcap.OpenOffline(pcapngFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()
	
	// 创建一个包解析器
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	
	// 源 IP 和源端口
	sourceIP := net.ParseIP("198.144.3.88")
	sourcePort := layers.TCPPort(6703)
	
	// 迭代处理每个包
	for packet := range packetSource.Packets() {
		// 解析包的各个层级
		ethLayer := packet.Layer(layers.LayerTypeEthernet)
		ipLayer := packet.Layer(layers.LayerTypeIPv4)
		tcpLayer := packet.Layer(layers.LayerTypeTCP)
		if ethLayer == nil || ipLayer == nil || tcpLayer == nil {
			log.Println("....continue")
			continue
		}
		
		// 解析 IP 和 TCP 层的数据
		ip, _ := ipLayer.(*layers.IPv4)
		tcp, _ := tcpLayer.(*layers.TCP)
		
		// 检查源 IP 和源端口是否匹配
		if ip.SrcIP.Equal(sourceIP) && tcp.SrcPort == sourcePort {
			// 打印数据包信息
			fmt.Printf(">>> Packet from %s:%d to %s:%d\n", ip.SrcIP, tcp.SrcPort, ip.DstIP, tcp.DstPort)
			fmt.Printf("> filename: %s\n", fileName)
			fmt.Printf("> Payload string: %s\n", string(tcp.Payload))
			fmt.Printf("> Payload hex: %x\n", tcp.Payload)
			
			// 构造响应数据包
			/*responsePacket := gopacket.NewPacket([]byte{}, layers.LayerTypeTCP, gopacket.Default)
			tcpLayer := responsePacket.Layer(layers.LayerTypeTCP).(*layers.TCP)
			tcpLayer.SrcPort = tcp.DstPort
			tcpLayer.DstPort = tcp.SrcPort
			tcpLayer.Seq = tcp.Ack
			tcpLayer.Ack = tcp.Seq + uint32(len(tcp.Payload))
			tcpLayer.ACK = true
			tcpLayer.RST = true
			
			// 发送响应数据包
			if err := handle.WritePacketData(responsePacket.Data()); err != nil {
				log.Println("Failed to send response packet:", err)
			} else {
				fmt.Println("Sent response packet")
			}*/
		}
	}
}
