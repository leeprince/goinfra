package main

import (
	"fmt"
	"github.com/globalsign/mgo"
	"log"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/10/9 16:07
 * @Desc:
 */

const (
	// 注意：27017后面是斜杠后跟着问号！
	mongoUri = "mongodb://mongoadmin:TB5i9K2jD1SAasdr@10.21.32.14:27017/?connect=direct"

	Database   = "ptest-db"
	Collection = "ptest-col"
)

// 结构体必须定义 bson 标签
type dataStruct struct {
	Name  string  `bson:"name" json:"name"`
	Value float64 `bson:"value" json:"value"`
}

var mongoDBConfig = MongoDBConfig{
	Uri:   mongoUri,
	Debug: true,
}

type MongoDBConfig struct {
	Uri   string
	Debug bool
}

var mongoClient *mgo.Session

func initMongoDBClient(config MongoDBConfig) {
	var err error
	mongoClient, err = mgo.Dial(config.Uri)
	if err != nil {
		panic(err)
	}

	if config.Debug {
		mgo.SetDebug(true)           // 设置DEBUG模式
		mgo.SetLogger(new(MongoLog)) // 设置日志.
	}
}

// 实现 mongo.Logger 的接口
type MongoLog struct {
}

func (MongoLog) Output(callDepth int, s string) error {
	log.SetFlags(log.Lshortfile)
	return log.Output(callDepth, s)
}

func main() {
	initMongoDBClient(mongoDBConfig)
}

func Insert() {
	insertData := dataStruct{
		Name:  "prince",
		Value: 2314.520,
	}
	err := mongoClient.DB(Database).C(Collection).Insert(&insertData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("insertData: %+v\n", insertData)
}
