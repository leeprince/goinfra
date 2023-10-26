package dao

import (
	"context"
	"github.com/leeprince/goinfra/storage/mongodb/mongogodrivertest/gd/app"
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

type RiskControlDao struct {
	ctx            context.Context
	logID          string
	riskControlCol *mongo.Collection
}

func NewRiskControlDao(ctx context.Context, logID string) *RiskControlDao {
	riskControl := model.RiskControl{}
	return &RiskControlDao{
		ctx:            ctx,
		logID:          logID,
		riskControlCol: app.LogMongoClient.Database(riskControl.DataBaseName()).Collection(riskControl.CollectionName()),
	}
}

func (r *RiskControlDao) Insert(data *model.RiskControl) (id primitive.ObjectID, err error) {
	res, err := r.riskControlCol.InsertOne(r.ctx, data)
	if err != nil {
		return
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

func (r *RiskControlDao) UpdateOneOfNotice(id primitive.ObjectID, isNotice bool, noticeFailReason string) (err error) {
	ctime := time.Now()

	// 不存在时是否插入
	opts := options.Update().SetUpsert(false)

	// 查询条件和设置的值可以参考插入和查询时使用的结构体方式
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{
		{model.RiskControlField.IsNotice, isNotice},
		{model.RiskControlField.NoticeFailReason, noticeFailReason},
		{model.RiskControlField.UpdatedAt, ctime.Unix()},
	}}}
	_, err = r.riskControlCol.UpdateOne(context.Background(), filter, update, opts)

	return
}
