package main

import "testing"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/10/9 16:22
 * @Desc:
 */

func init() {
	// 初始化MongoDB客户端
	initMongoDBClient(mongoDBConfig)
}

func TestInsert(t *testing.T) {
	Insert()
}
