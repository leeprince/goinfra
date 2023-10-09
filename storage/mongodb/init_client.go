package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/10/9 16:32
 * @Desc:
 */

type MongoDBConfig struct {
	Uri string
}

// InitMongoDBClient 初始化MongoDB客户端
func InitMongoDBClient(config MongoDBConfig) (mongoClient *mongo.Client, err error) {
	// 连接  mongoDB 服务端
	ctx := context.Background()
	mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(config.Uri))
	// 使用全局变量，故不直接在 defer 中关闭连接
	/*defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()*/

	if err != nil {
		return
	}

	// 验证连接
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		return
	}

	return
}
