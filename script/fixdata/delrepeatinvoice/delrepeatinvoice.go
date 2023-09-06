package delrepeatinvoice

import (
	"fmt"
	"github.com/spf13/cobra"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/9/5 17:25
 * @Desc:
 */
type option struct {
	Env string

	OrgId          string
	EId            string
	CUserId        string
	OrderSn        string
	AfterUpdatedAt string
	IsUpdate       bool
}

func New() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "delrepeatinvoice",
		Short: "软删除重复且未删除状态的历史发票（保留最新）",
		RunE:  findReatInvoiceAndDelete,
	}

	rootCmd.PersistentFlags().Bool("is_update", false, "是否执行更新sql")
	rootCmd.PersistentFlags().String("org_id", "", "租户ID")
	rootCmd.PersistentFlags().String("eid", "", "企业ID")
	rootCmd.PersistentFlags().String("c_user_id", "", "c端用户ID")
	rootCmd.PersistentFlags().String("order_sn", "", "发票order_sn")
	rootCmd.PersistentFlags().String("after_updated_at", "2023-08-26 00:00:00", "删除指定时间到现在,时间格式：2023-08-26 00:00:00")
	return rootCmd
}

func initFlags(c *cobra.Command) (opt *option, err error) {
	env, err := c.Parent().PersistentFlags().GetString("env")
	if err != nil {
		fmt.Println("env err:", err)
		return
	}
	isUpdate, err := c.PersistentFlags().GetBool("is_update")
	if err != nil {
		fmt.Println("after_updated_at err:", err)
		return
	}
	orgId, err := c.PersistentFlags().GetString("org_id")
	if err != nil {
		fmt.Println("org_id err:", err)
		return
	}
	eid, err := c.PersistentFlags().GetString("eid")
	if err != nil {
		fmt.Println("eid err:", err)
		return
	}
	cUserId, err := c.PersistentFlags().GetString("c_user_id")
	if err != nil {
		fmt.Println("c_user_id err:", err)
		return
	}
	orderSn, err := c.PersistentFlags().GetString("order_sn")
	if err != nil {
		fmt.Println("order_sn err:", err)
		return
	}
	afterUpdatedAt, err := c.PersistentFlags().GetString("after_updated_at")
	if err != nil {
		fmt.Println("after_updated_at err:", err)
		return
	}

	opt = &option{
		Env:            env,
		IsUpdate:       isUpdate,
		OrgId:          orgId,
		EId:            eid,
		CUserId:        cUserId,
		OrderSn:        orderSn,
		AfterUpdatedAt: afterUpdatedAt,
	}

	return
}

func findReatInvoiceAndDelete(c *cobra.Command, _ []string) (err error) {
	opt, err := initFlags(c)
	if err != nil {
		fmt.Println("initFlags err:", err)
		return
	}
	fmt.Printf("opt:%+v \n", opt)

	// 检查 option
	if opt.OrgId == "" &&
		opt.EId == "" &&
		opt.CUserId == "" &&
		opt.OrderSn == "" &&
		opt.AfterUpdatedAt == "" {
		fmt.Println("必须填写其中一个条件(org_id,eid,c_user_id,order_sn,after_updated_at)")
		return
	}

	return
}
