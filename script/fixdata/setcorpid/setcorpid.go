package setcorpid

import (
	"fmt"
	"github.com/leeprince/goinfra/script/fixdata/initconfig"
	"github.com/spf13/cobra"
	"strings"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/11 17:09
 * @Desc:
 */

type option struct {
	Env string

	UserIdList []string
	OrgId      string
	OrderSn    string
	CropId     string
}

func New() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "setcorpid",
		Short: "根据用户ID查找(指定org_id&order_sn)&corp_id=''的记录，并设置指定corp_id",
		RunE:  findUserDcacheDataAndSetCorpId,
	}
	rootCmd.PersistentFlags().String("user_id", "", "c端用户ID:多个用户时使用英文逗号,分开")
	rootCmd.PersistentFlags().String("org_id", "", "c端用户ID下指定的org_id")
	rootCmd.PersistentFlags().String("order_sn", "", "c端用户ID下指定的order_sn")
	rootCmd.PersistentFlags().String("corp_id", "", "需要设置的corp_id")
	return rootCmd
}

func initFlags(c *cobra.Command) (opt *option, err error) {
	env, err := c.Root().PersistentFlags().GetString("env")
	if err != nil {
		fmt.Println("env err:", err)
		return
	}
	userIdListString, err := c.PersistentFlags().GetString("user_id")
	if err != nil {
		fmt.Println("UserId err:", err)
		return
	}
	orgId, err := c.PersistentFlags().GetString("org_id")
	if err != nil {
		fmt.Println("userId err:", err)
		return
	}
	orderSn, err := c.PersistentFlags().GetString("order_sn")
	if err != nil {
		fmt.Println("orderSn err:", err)
		return
	}
	corpId, err := c.PersistentFlags().GetString("corp_id")
	if err != nil {
		fmt.Println("orgId err:", err)
		return
	}

	opt = &option{
		Env:        env,
		UserIdList: strings.Split(userIdListString, ","),
		OrgId:      orgId,
		OrderSn:    orderSn,
		CropId:     corpId,
	}

	fmt.Printf("opt:%+v \n", opt)

	return
}

func findUserDcacheDataAndSetCorpId(c *cobra.Command, _ []string) (err error) {
	opt, err := initFlags(c)
	if err != nil {
		fmt.Println("initFlags err:", err)
		return
	}

	err = initconfig.Init(opt.Env)
	if err != nil {
		return
	}

	if len(opt.UserIdList) <= 0 {
		fmt.Println("opt.UserIdList empty")
		return
	}
	if opt.OrgId == "" {
		fmt.Println("opt.OrgId empty & OrgId 为空时只能是0字符串")
		return
	}

	return
}
