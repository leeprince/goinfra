package main

import (
	"fmt"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/10/9 14:50
 * @Desc:
 */

func init() {
	// 初始化MongoDB客户端
	initMongoDBClient()
}

func TestInsertOne(t *testing.T) {
	InsertOne()
}

func TestFind(t *testing.T) {
	Find()

	fmt.Println("--------------------------------------------------------------------")

	FindOne()
}

func TestUpdateOne(t *testing.T) {
	UpdateOne()
}
