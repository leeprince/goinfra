package idutil

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/5/8 20:52
 * @Desc:	自定义唯一ID
 * 				性能：UniqIDV3 > UniqIDV2 > UniqIDV1
 */

func UniqIDV1() string {
	now := time.Now().UnixNano()
	rand.Seed(now)             // 设置随机种子，使用当前时间的纳秒数
	randInt := rand.Intn(1000) // 生成0到999之间的随机整数

	// %016x: 用0填充最小宽度16，十六进制格式
	return fmt.Sprintf("%016x", now+int64(randInt))
}

func UniqIDV2() string {
	uuid := UUID()
	idArr := strings.Split(uuid, "-")
	return idArr[0] + idArr[1] + idArr[2]
}

func UniqIDV3() string {
	return fmt.Sprintf("%016x", NewSnowflake(maxWorker).NextId())
}