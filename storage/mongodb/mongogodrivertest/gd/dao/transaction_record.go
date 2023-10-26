package dao

import (
	"context"
	"github.com/leeprince/goinfra/storage/mongodb/mongogodrivertest/gd/app"
	"github.com/leeprince/goinfra/storage/mongodb/mongogodrivertest/gd/constants"
	"github.com/leeprince/goinfra/storage/mongodb/mongogodrivertest/gd/dao/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/10/20 16:44
 * @Desc:
 */

type TransactionRecordDao struct {
	ctx             context.Context
	logID           string
	operationLogCol *mongo.Collection
}

func NewTransactionRecordDao(ctx context.Context, logID string) *TransactionRecordDao {
	operationLog := model.TransactionRecord{}
	return &TransactionRecordDao{
		ctx:             ctx,
		logID:           logID,
		operationLogCol: app.LogMongoClient.Database(operationLog.DataBaseName()).Collection(operationLog.CollectionName()),
	}
}

func (r *TransactionRecordDao) Insert(data *model.TransactionRecord) (id primitive.ObjectID, err error) {
	res, err := r.operationLogCol.InsertOne(r.ctx, data)
	if err != nil {
		return
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

func (r *TransactionRecordDao) FindOne(openOrgId, transactionSn string, transactionStatus constants.TransactionStatus) (result *model.QueryTransactionRecord, err error) {
	/*通过bson.D作为查询条件，并查询到结构体*/
	result = &model.QueryTransactionRecord{}
	filter := bson.D{
		{model.TransactionRecordField.OpenOrgId, openOrgId},
		{model.TransactionRecordField.TransactionSn, transactionSn},
		{model.TransactionRecordField.TransactionStatus, transactionStatus},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = r.operationLogCol.FindOne(ctx, filter).Decode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = nil
			return
		}
	}

	return
}

func (r *TransactionRecordDao) UpdateOneOfStatus(id primitive.ObjectID, status constants.TransactionStatus) (err error) {
	// 不存在时是否插入
	opts := options.Update().SetUpsert(false)

	// 查询条件和设置的值可以参考插入和查询时使用的结构体方式
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{model.TransactionRecordField.TransactionStatus, status}}}}
	_, err = r.operationLogCol.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		return
	}

	return
}
