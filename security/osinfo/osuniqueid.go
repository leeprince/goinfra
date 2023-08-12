package osinfo

import (
	"fmt"
	"github.com/elastic/go-sysinfo"
	"strings"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/12 00:46
 * @Desc:
 */

func GetOsUniqueId() (mark string, err error) {
	// 获取系统信息
	sysInfo, err := sysinfo.Host()
	if err != nil {
		return
	}
	
	info := sysInfo.Info()
	uniqueId := info.UniqueID
	kernelVersion := info.KernelVersion
	hostname := info.Hostname
	macsString := strings.Join(info.MACs, ",")
	
	mark = fmt.Sprintf("%s;%s;%s;%s;", uniqueId, kernelVersion, hostname, macsString)
	
	return
}
