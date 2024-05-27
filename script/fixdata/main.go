package main

import (
	"fmt"
	"github.com/leeprince/goinfra/script/fixdata/delrepeatinvoice"
	"github.com/leeprince/goinfra/script/fixdata/setcorpid"
	"github.com/leeprince/goinfra/script/fixdata/version"
	"github.com/spf13/cobra"
	"runtime/debug"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/11 17:06
 * @Desc:
 */

/*
使用说明：
fixdata 不用写

################## 帮助命令
go run main.go help

go run main.go help setcorpid
go run main.go help delrepeatinvoice

################## setcorpid
# 测试
go run main.go setcorpid --user_id 20201711 --org_id 0 --order_sn="6669878947741438535" --corp_id=""
go run main.go setcorpid --user_id 20201711 --org_id 0 --order_sn="" --corp_id=""
# 生产
go run main.go setcorpid --user_id 445168 --org_id 488870 --order_sn="7086993366806041064" --corp_id="1609813763522568298"
go run main.go setcorpid --user_id 17004576 --org_id 488870 --order_sn="" --corp_id="1609813763522568298"
go run main.go setcorpid --user_id 15896274 --org_id 488870 --order_sn="" --corp_id="1609813763522568298"
go run main.go setcorpid --user_id 17004576,15896274 --org_id 488870 --order_sn="" --corp_id="1609813763522568298"
go run main.go setcorpid --user_id 17004576,16428581 --org_id 488870 --order_sn="" --corp_id="1609813763522568298"
go run main.go setcorpid --user_id 16007312,23935025,16702494,15896309,24621220 --org_id 488870 --order_sn="" --corp_id="1609813763522568298"

################## delrepeatinvoice
# 测试
go run main.go delrepeatinvoice -e test
go run main.go delrepeatinvoice --env test

# 注意：--after_updated_at 参数是有空格的，所以flag必须要用括号""包起来；--is_update 是布尔类型，所以flag也必须适用等于号（=）跟着键值；
go run main.go delrepeatinvoice -e test --after_updated_at "2023-08-19 00:00:00"
go run main.go delrepeatinvoice -e test --after_updated_at="2023-08-19 00:00:00"
go run main.go delrepeatinvoice -e test --after_updated_at="2023-08-19 00:00:00" --is_update=false
go run main.go delrepeatinvoice -e test --after_updated_at="2023-08-19 00:00:00" --is_update=true
go run main.go delrepeatinvoice -e test --after_updated_at="2023-08-19 00:00:00" --org_id 1
go run main.go delrepeatinvoice -e test --after_updated_at="2023-08-19 00:00:00" --eid xxx
go run main.go delrepeatinvoice -e test --after_updated_at="2023-08-19 00:00:00" --order_sn xxxz
go run main.go delrepeatinvoice -e test --after_updated_at="2023-08-19 00:00:00" --c_user_id 1


go run main.go delrepeatinvoice -e test

# 注意：--is_update 参数是是bool值，所以必须要用等号=的方式传参
go run main.go delrepeatinvoice -e test --is_update=true


*/

/*
交叉编译执行：
GOOS=windows GOARCH=amd64 go build -o fixdata.exe main.go

GOOS=linux GOARCH=amd64 go build -o fixdata_linux_64 main.go
*/

func main() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("main panic: %v stack: %s", p, string(debug.Stack()))
		}
	}()

	if err := New().Execute(); err != nil {
		panic(err)
	}
}

func New() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "fixdata",
		Short: "修复 mysql/dcache 数据脚本",
		Run: func(cmd *cobra.Command, args []string) {
			// 执行根命令时可以做的事情
			// ...
		},
	}

	rootCmd.PersistentFlags().StringP("env", "e", "test", "environment: dev, test, prod")

	rootCmd.AddCommand(version.New())
	rootCmd.AddCommand(setcorpid.New())
	rootCmd.AddCommand(delrepeatinvoice.New())

	return rootCmd
}
