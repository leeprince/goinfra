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
	SerialDataHeadLen = 4
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
			buflen := make([]byte, SerialDataHeadLen) // 调整缓冲区大小
			n, err := port.Read(buflen)
			log.Println("读取到长度:" + strconv.Itoa(n))
			if err != nil {
				log.Println(err)
				return
			}
			if n == 0 {
				continue
			}
			for n < SerialDataHeadLen {
				needCount := SerialDataHeadLen - n
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
			if n != SerialDataHeadLen {
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
			
			// 报文请求体内容的长度
			bodyLen := len(line)
			
			// 请求内容:报文请求头的长度+报文请求体内容的长度
			reqByte := make([]byte, SerialDataHeadLen+bodyLen)
			
			// 设置报文的请求头的内容
			bodyLenByte := make([]byte, SerialDataHeadLen)
			binary.BigEndian.PutUint32(bodyLenByte, uint32(bodyLen))
			copy(reqByte, bodyLenByte)
			
			// 设置报文的请求体的内容
			copy(reqByte[SerialDataHeadLen:], bodyByte)
			
			n, err := port.Write(reqByte)
			fmt.Println("n:", n)
			fmt.Println("err:", err)
			if err != nil {
				return
			}
			
			time.Sleep(3 * time.Second)
		}
	}()
	for {
	
	}
	
}
