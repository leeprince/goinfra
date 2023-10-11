package main

import (
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
}

func TestFindFilter(t *testing.T) {
	FindFilter()
}

func TestFindAndSort(t *testing.T) {
	FindAndSort()
}

func TestCount(t *testing.T) {
	Count()
}

func TestPage(t *testing.T) {
	Page()
}

func TestFindOne(t *testing.T) {
	FindOne()
}

func TestUpdateOne(t *testing.T) {
	UpdateOne()
}
