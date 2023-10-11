package dao

import (
	"context"
	"github.com/leeprince/goinfra/storage/mongodb/mongogodrivertest/gd/app"
	"github.com/leeprince/goinfra/storage/mongodb/mongogodrivertest/gd/dao/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/10/10 17:46
 * @Desc:
 */

type OperationLogDao struct {
	ctx             context.Context
	logID           string
	operationLogCol *mongo.Collection
}

func NewOperationLogDao(ctx context.Context, logID string) *OperationLogDao {
	operationLog := model.InsertOperationLogModel{}
	return &OperationLogDao{
		ctx:             ctx,
		logID:           logID,
		operationLogCol: app.LogMongoClient.Database(operationLog.DataBaseName()).Collection(operationLog.CollectionName()),
	}
}

func (r *OperationLogDao) Insert(data *model.InsertOperationLogModel) (err error) {
	_, err = r.operationLogCol.InsertOne(r.ctx, data)
	return
}

func (r *OperationLogDao) FindList(orgId int64, page, pageSize int64) (count int64, resp []*model.QueryOperationLogModel, err error) {
	filter := bson.D{{model.OperationLogField.OrgId, orgId}}

	count, err = r.operationLogCol.CountDocuments(r.ctx, filter)
	if err != nil {
		return
	}

	if count <= 0 {
		return
	}

	skip := (page - 1) * pageSize
	opt := options.Find().SetSort(bson.D{{model.OperationLogField.OperatedAt, -1}}).SetSkip(skip).SetLimit(pageSize)

	cur, err := r.operationLogCol.Find(r.ctx, filter, opt)
	if err != nil {
		return
	}
	defer cur.Close(r.ctx)

	resp = []*model.QueryOperationLogModel{}
	err = cur.All(r.ctx, &resp)
	if err != nil {
		return
	}
	return
}
