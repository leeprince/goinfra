package idutil

import (
	"github.com/google/uuid"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/16 上午10:28
 * @Desc:   UUID
 */

func UUID() string {
	return uuid.New().String()
}
