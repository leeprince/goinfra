package osinfo

import (
	"encoding/base64"
	"fmt"
	"github.com/elastic/go-sysinfo"
	"github.com/leeprince/goinfra/utils/arrayutil"
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
	hostname := info.Hostname
	uniqueId := info.UniqueID
	kernelVersion := info.KernelVersion
	
	mark = fmt.Sprintf("%s;%s;%s;", hostname, uniqueId, kernelVersion)
	
	return
}

func GetOsHostname() (hostname string, err error) {
	// 获取系统信息
	sysInfo, err := sysinfo.Host()
	if err != nil {
		return
	}
	
	info := sysInfo.Info()
	hostname = info.Hostname
	
	return
}

func GetOsUniqueIdBase64() (mark string, err error) {
	mark, err = GetOsUniqueId()
	if err != nil {
		return
	}
	
	mark = base64.StdEncoding.EncodeToString([]byte(mark))
	
	return
}

func CheckOsUniqueIdBase64(in string) (b bool) {
	mark, err := GetOsUniqueIdBase64()
	if err != nil {
		return
	}
	
	if mark == in {
		b = true
		return
	}
	
	return
}

func CheckOsUniqueIdBase64InArr(inArr []string) (b bool) {
	mark, err := GetOsUniqueIdBase64()
	if err != nil {
		return
	}
	
	if arrayutil.InString(mark, inArr) {
		b = true
		return
	}
	
	return
}
