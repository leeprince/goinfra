package ppostgres

import "time"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/6 下午11:39
 * @Desc:   常量定义
 */

const (
	MaxIdleConns    = 10
	MaxOpenConns    = 100
	ConnMaxLifetime = time.Hour
)
