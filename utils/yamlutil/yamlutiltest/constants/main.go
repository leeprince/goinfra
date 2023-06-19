package main

import (
	"embed"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/16 23:01
 * @Desc:
 */

//go:embed station.yaml
var stationsData embed.FS

type StationInfo struct {
	QuanPinMa string `yaml:"全拼码"`
	DianBaoMa string `yaml:"电报码"`
	JianPinMa string `yaml:"简拼码"`
}

var stations map[string]StationInfo

func init() {
	// 读取 YAML 文件
	/* stationsData.ReadFile("station.yaml") 配合 //go:embed station.yaml 使用
	//go:embed station.yaml
	var stationsData embed.FS
	*/
	data, err := stationsData.ReadFile("station.yaml")
	/* stationsData.ReadFile("configs/station.yaml") 配合 //go:embed configs/station.yaml 使用
	//go:embed configs/station.yaml
	var stationsData embed.FS
	*/
	// data, err := stationsData.ReadFile("configs/station.yaml")
	if err != nil {
		log.Fatalf("failed to read file %v", err)
	}

	// 解析 YAML 文件
	err = yaml.Unmarshal(data, &stations)
	if err != nil {
		log.Fatalf("failed to unmarshal data: %v", err)
	}
}

func main() {
	// 输出结果
	fmt.Printf("一面坡北：全拼码：%s，电报码：%s，简拼码：%s\n", stations["一面坡北"].QuanPinMa, stations["一面坡北"].DianBaoMa, stations["一面坡北"].JianPinMa)
	fmt.Printf("一面山：全拼码：%s，电报码：%s，简拼码：%s\n", stations["一面山"].QuanPinMa, stations["一面山"].DianBaoMa, stations["一面山"].JianPinMa)
	fmt.Printf("七台河：全拼码：%s，电报码：%s，简拼码：%s\n", stations["七台河"].QuanPinMa, stations["七台河"].DianBaoMa, stations["七台河"].JianPinMa)
	fmt.Printf("不存在的：全拼码：%s，电报码：%s，简拼码：%s\n", stations["不存在的"].QuanPinMa, stations["不存在的"].DianBaoMa, stations["不存在的"].JianPinMa)
}
