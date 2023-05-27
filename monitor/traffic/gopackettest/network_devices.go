package main

import (
	"fmt"
	"github.com/google/gopacket/pcap"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/5/4 22:29
 * @Desc:
 */

func NetworkDevices() {
	// 获取所有网络接口信息
	devices, err := pcap.FindAllDevs()
	if err != nil {
		panic(err)
	}

	// 打印每个网络接口的信息
	for _, device := range devices {
		fmt.Println("Name:", device.Name)
		fmt.Println("Description:", device.Description)
		fmt.Println("Addresses:")
		for _, address := range device.Addresses {
			fmt.Println("- IP:", address.IP)
			fmt.Println("  Netmask:", address.Netmask)
		}
	}
}
