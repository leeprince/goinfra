package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"github.com/tarm/serial"
	"log"
	"os"
	"strconv"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/9/2 19:15
 * @Desc:
 */

const (
	dataHeadLen = 4
)

func main() {
	
	// 配置串口参数
	c := &serial.Config{
		Name: "COM1",
		Baud: 9600,
	}
	
	// 打开串口
	port, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()
	
	// 接收数据
	go func() {
		for {
			buflen := make([]byte, dataHeadLen) // 调整缓冲区大小
			n, err := port.Read(buflen)
			log.Println("读取到长度:" + strconv.Itoa(n))
			if err != nil {
				log.Println(err)
				return
			}
			if n == 0 {
				continue
			}
			for n < dataHeadLen {
				needCount := dataHeadLen - n
				readBody := make([]byte, needCount)
				count, err := port.Read(readBody)
				if err != nil {
					log.Println(err)
					return
				}
				if count == 0 {
					continue
				}
				copy(buflen[n:], readBody)
				n += count
			}
			if n != dataHeadLen {
				log.Println("读取到字节数不对")
			}
			// 将字节数组转换为 uint32
			bodyLen := binary.BigEndian.Uint32(buflen)
			
			bufBody := make([]byte, bodyLen)
			readCount, err := port.Read(bufBody)
			if err != nil {
				log.Println(err)
				return
			}
			if readCount == 0 {
				continue
			}
			for readCount < int(bodyLen) {
				needCount := int(bodyLen) - readCount
				readBody := make([]byte, needCount)
				count, err := port.Read(readBody)
				if err != nil {
					log.Println(err)
					return
				}
				if count == 0 {
					continue
				}
				copy(bufBody[readCount:], readBody)
				readCount += count
			}
			if readCount != int(bodyLen) {
				log.Println("读取到字节数不对")
			}
			log.Println("body内容:" + string(bufBody))
			
		}
	}()
	
	// 发送数据
	go func() {
		file, err := os.Open("order.txt")
		if err != nil {
			fmt.Println("无法打开文件:", err)
			return
		}
		defer file.Close()
		
		scanner := bufio.NewScanner(file)
		
		for scanner.Scan() {
			line := scanner.Text()
			bodyByte := []byte(line)
			bodyLen := len(line)
			
			// 请求内容:dataHeadLen(存放报文长度)+报文内容
			reqByte := make([]byte, dataHeadLen+bodyLen)
			bodyLenByte := make([]byte, dataHeadLen)
			
			binary.BigEndian.PutUint32(bodyLenByte, uint32(bodyLen))
			copy(reqByte, bodyLenByte)
			
			copy(reqByte[dataHeadLen:], bodyByte)
			
			port.Write(reqByte)
			time.Sleep(3 * time.Second)
		}
	}()
	for {
	
	}
	
}
