package main

import (
	"github.com/leeprince/goinfra/script/fixdatatest/setcorpid"
	"github.com/spf13/cobra"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/11 17:06
 * @Desc:
 */

/*
使用说明：
go run main.go --help
go run main.go setcorpid --help

注意：fixdata是根命令不用写，也不能写

例子：
go run main.go setcorpid --user_id 20201711 --org_id 0 --order_sn 6669878947741438535 --corp_id=""
*/
func main() {
	if err := New().Execute(); err != nil {
		panic(err)
	}
}

func init() {
	// 配置

}

func New() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "fixdata",
		Short: "修复数据",
	}

	rootCmd.PersistentFlags().StringP("env", "e", "dev", "environment: dev, test, prod")
	rootCmd.AddCommand(setcorpid.New())

	return rootCmd
}
