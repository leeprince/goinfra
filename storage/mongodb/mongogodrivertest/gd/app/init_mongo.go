package app

import (
	"fmt"
	"github.com/leeprince/goinfra/storage/mongodb"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/sync/errgroup"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/10/9 16:37
 * @Desc:
 */

var (
	LogMongoClient *mongo.Client
)

const (
	// 注意：27017后面是斜杠后跟着问号！
	mongoUri = "mongodb://mongoadmin:TB5i9K2jD1SAasdr@10.21.32.14:27017/?connect=direct"
)

// InitMongoDBClient 初始化MongoDB客户端
func InitMongoDBClient() {
	errg := errgroup.Group{}
	errg.Go(func() (err error) {
		// 连接  mongoDB 服务端
		config := mongodb.MongoDBConfig{
			Uri: mongoUri,
		}
		LogMongoClient, err = mongodb.InitMongoDBClient(config)
		if err != nil {
			return errors.New("InitMongoDBClient Connect Log 失败" + err.Error())
		}

		return
	})

	err := errg.Wait()
	if err != nil {
		panic("InitMongoDBClient 初始化失败" + err.Error())
	}

	fmt.Println("InitMongoDBClient end")
}
