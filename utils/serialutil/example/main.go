package main

import (
	"github.com/tarm/serial"
	"log"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/9/2 19:16
 * @Desc:
 */

func main() {
	// 配置串口参数
	config := &serial.Config{
		Name:        "/dev/tty.Bluetooth-Incoming-Port", // 串口设备名称，根据实际情况修改
		Baud:        9600,                               // 波特率，根据实际情况修改
		ReadTimeout: time.Second,                        // 读取超时时间
	}
	
	// 打开串口
	port, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal("serial.OpenPort err:", err)
	}
	defer port.Close()
	log.Println("serial.OpenPort success")
	
	// 接收数据
	go func() {
		for {
			buf := make([]byte, 128)
			n, err := port.Read(buf)
			if err != nil {
				if err.Error() == "EOF" {
					// log.Println("port.Read error Is EOF > continue...")
					continue
				}
				log.Printf("port.Read err1:%+v \n", err)
				log.Printf("port.Read err2:%#v \n", err)
			}
			log.Printf("port.Read data:%q\n", buf[:n])
			log.Printf("port.Read data:%v\n", string(buf[:n]))
		}
	}()
	
	time.Sleep(time.Second * 1)
	
	// 发送数据
	data := []byte("Hello, RS232!")
	_, err = port.Write(data)
	if err != nil {
		log.Fatal("port.Write err:", err)
	}
	log.Println("port.Write success data:", string(data))
	
	select {}
}
