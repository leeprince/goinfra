package wiresharkfile

import (
	"fmt"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/8 20:16
 * @Desc:
 */

func TestMonitorPcapngFileTcp(t *testing.T) {
	type args struct {
		pcapngFilePath string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "#-c7002-01-01A-一等-995.pcapng",
			args: args{
				pcapngFilePath: "/Users/leeprince/www/go/goinfra/monitor/traffic/gopackettest/wiresharkfile/file/c7002-01-01A-一等-995.pcapng",
			},
		},
		{
			name: "#-c7002-05-02A-二等-795.pcapng",
			args: args{
				pcapngFilePath: "/Users/leeprince/www/go/goinfra/monitor/traffic/gopackettest/wiresharkfile/file/c7002-05-02A-二等-795.pcapng",
			},
		},
		{
			name: "#-c7006-01-01A-01-01C-一等-995.pcapng",
			args: args{
				pcapngFilePath: "/Users/leeprince/www/go/goinfra/monitor/traffic/gopackettest/wiresharkfile/file/c7006-01-01A-01-01C-一等-995.pcapng",
			},
		},
		{
			name: "#-c7006-01-01A-一等-995.pcapng",
			args: args{
				pcapngFilePath: "/Users/leeprince/www/go/goinfra/monitor/traffic/gopackettest/wiresharkfile/file/c7006-01-01A-一等-995.pcapng",
			},
		},
		{
			name: "#-c7006-03-01A-03-01C-二等-795.pcapng",
			args: args{
				pcapngFilePath: "/Users/leeprince/www/go/goinfra/monitor/traffic/gopackettest/wiresharkfile/file/c7006-03-01A-03-01C-二等-795.pcapng",
			},
		},
		{
			name: "#-c7006-04-01A-二等-795.pcapng",
			args: args{
				pcapngFilePath: "/Users/leeprince/www/go/goinfra/monitor/traffic/gopackettest/wiresharkfile/file/c7006-04-01A-二等-795.pcapng",
			},
		},
		{
			name: "#-k838-13-21上铺-硬卧-1290.pcapng",
			args: args{
				pcapngFilePath: "/Users/leeprince/www/go/goinfra/monitor/traffic/gopackettest/wiresharkfile/file/k838-13-21上铺-硬卧-1290.pcapng",
			},
		},
		{
			name: "#-k838-14-09下铺-1360-14-09上铺-1290-硬卧.pcapng",
			args: args{
				pcapngFilePath: "/Users/leeprince/www/go/goinfra/monitor/traffic/gopackettest/wiresharkfile/file/k838-14-09下铺-1360-14-09上铺-1290-硬卧.pcapng",
			},
		},
		{
			name: "#-k838-14-09下铺-1360-硬卧.pcapng",
			args: args{
				pcapngFilePath: "/Users/leeprince/www/go/goinfra/monitor/traffic/gopackettest/wiresharkfile/file/k838-14-09下铺-1360-硬卧.pcapng",
			},
		},
		{
			name: "#-k838-15-08中铺-1320-硬卧.pcapng",
			args: args{
				pcapngFilePath: "/Users/leeprince/www/go/goinfra/monitor/traffic/gopackettest/wiresharkfile/file/k838-15-08中铺-1320-硬卧.pcapng",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("------------------------------------------------------------\n\n\n")
			
			MonitorPcapngFileTcp(tt.args.pcapngFilePath)
		})
	}
}
