package app

import (
	"context"
	"fmt"
	"github.com/leeprince/goinfra/storage/mongodb/mongogodrivertest/gd/config"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/sync/errgroup"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/10/9 16:37
 * @Desc:	初始化mongDB客户端
 */

var (
	LogMongoClient *mongo.Client
)

// InitMongoDBClient 初始化MongoDB客户端
func InitMongoDBClient() {
	errg := errgroup.Group{}
	errg.Go(func() (err error) {
		// 连接  mongoDB 服务端
		ctx := context.Background()
		LogMongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(config.Config.MongoDB.Log.Uri))
		if err != nil {
			return errors.New("InitMongoDBClient Connect Log 失败" + err.Error())
		}

		// 验证连接
		err = LogMongoClient.Ping(ctx, readpref.Primary())
		if err != nil {
			return errors.New("InitMongoDBClient Ping Log 失败" + err.Error())
		}
		return
	})

	err := errg.Wait()
	if err != nil {
		panic("InitMongoDBClient 初始化失败" + err.Error())
	}

	fmt.Println("InitMongoDBClient end")
}
