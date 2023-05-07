# 流量监控
---


# 监听流量的库
| 库名称        |  特点 |    优点 |    缺点|
|------------|  --- |   --- |   ---|
| go-netmon  |    基于 libpcap 的网络监控库，支持捕获和分析本地网络流量。 |  简单易用，支持多种议。     |功能相对较少，不支持高级过滤条件。|
| gopacket   | 用于网络数据包分的库，可以捕获和解析本地网络流量，支持多种协议。|   功能强大，支持多种协议和高级过滤条件 |    学习成本较高，使用较为复杂。|
| tcpdump    |  基于 libpcap 的命令行具，可以捕获和分析本网络流量，支持多种过滤条件。 |   功能强大，支持多种过滤条件。 |    使用命令行操作，不够直观，需要一定的命令行基础。|
| tshark     |    基于 Wireshark 的命令行工具，可以捕获和分析本地网络流量，支持多种过滤条件和输出格式。 |  功能强大，支持多种过滤条件和输出。 | 使用命令行操作，不够直观，需要一定的命令行基础。|
| prometheus |   用于监控和度量的系统，可以通过客户端库和 exporter 监控本地网络流量其他系统指标。 | 功能强大，支持多种指标监控。  |学习成本较高，需要一定的系统管理和监控经验。|

综上所述，不同的库适用于不同的场景，需要根据实际需求选择合适的库进行使用。
如果需要简单的网络流量监控，可以选择 go-netmon 或 tcpdump；
如果需要更为强大的功能和高过滤条件，可以选择 gopacket 或 tshark；
如果需要监控和度量多种指标，可以选择 prometheus。


# 确定网卡设备
## mac
在 Mac 中，可以通过终端命令 ifconfig 来查看当前系统中的网络接口信息。该命令会列出所有的网络接，包括以太网、Wi-Fi、蓝牙等。每个网络接口都有一个名称，通常以 en 开头，例如 en0、1 等。

在确定应该传递给 pcap.OpenLive 函数的 device 参数时，可以根据需要监听的网络接口名称来确定。例如，如果需要监听 Wi-Fi 接口，可以将 device 参数设置为 en0。如果需要监听以太网接口，可以将 device 参数设置为 en1。

需要注意的是，在 Mac 中，需要使用管理员权限运行才能够监听网络接口。可以使用 sudo 命令来运行程序，例如：

sudo ./my-program

在程序中，可以使用 pcap.FindAllDevs 函数来获取当前系统中的所有网络接口信息，包括名称、描述、IP 地址等。例如：

```

package main

import (
    "fmt"
    "github.com/google/gopacket/pcap"
)

func main() {
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
```

在上面的代码中，我们使用 pcap.FindAllDevs 函数获取所有网络接口信息，并打印出每个网络口的名称、描述和 IP 地址等信息。可以根据这些信息来确定应该传递给 pcap.OpenLive 函数的 device 参数

# 测试
## 有权限地运行`gopackettest/main.go`
```
sudo go run main.go
```

## 模拟请求
```
curl example.com
```