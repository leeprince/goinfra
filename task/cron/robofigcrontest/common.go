package robofigcrontest

import (
	"github.com/robfig/cron/v3"
	"log"
	"os"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/21 10:31
 * @Desc:
 */

func NewWithSeconds() *cron.Cron {
	// 支持秒级定时表达式
	/*secondParser := cron.NewParser(cron.Second |
		cron.Minute |
		cron.Hour |
		cron.Dom |
		cron.Month |
		cron.DowOptional |
		cron.Descriptor)
	cronInstance := cron.New(cron.WithParser(secondParser))*/

	// 支持秒级定时表达式: 简化写法
	cronInstance := cron.New(cron.WithSeconds())

	return cronInstance
}

// NewWithSecondsWithChain 支持上次任务未执行完，下次任务不启动: 默认上次任务没运行完，下次任务依然会运行（任务运行在goroutine里相互不干扰）
func NewWithSecondsWithChain() *cron.Cron {
	cronInstance := cron.New(
		cron.WithSeconds(),
		cron.WithChain(
			cron.SkipIfStillRunning(
				cron.VerbosePrintfLogger(
					log.New(os.Stdout, "cron.SkipIfStillRunning", log.LstdFlags)))),
	)

	return cronInstance
}
