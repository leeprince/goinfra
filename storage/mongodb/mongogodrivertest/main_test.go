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

func TestInsertOneByStruct(t *testing.T) {
	InsertOneByStruct()
}

func TestFind(t *testing.T) {
	Find()
}

func TestFindOneById(t *testing.T) {
	FindOneById()
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

func TestFindOneOfMoreFilter(t *testing.T) {
	FindOneOfMoreFilter()
}

func TestUpdateOne(t *testing.T) {
	UpdateOne()
}

func TestInsertOneAndUpdateOne(t *testing.T) {
	InsertOneAndUpdateOne()
}
