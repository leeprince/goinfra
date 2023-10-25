package main

import (
	"context"
	"github.com/leeprince/goinfra/storage/mongodb/mongogodrivertest/gd/app"
	dao2 "github.com/leeprince/goinfra/storage/mongodb/mongogodrivertest/gd/dao"
	"github.com/leeprince/goinfra/storage/mongodb/mongogodrivertest/gd/dao/model"
	"github.com/leeprince/goinfra/utils/dumputil"
	"github.com/leeprince/goinfra/utils/idutil"
	"testing"
	"time"
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

func TestFindOneOfMoreFilter(t *testing.T) {
	FindOneOfMoreFilter()
}

func TestUpdateOne(t *testing.T) {
	UpdateOne()
}

/* ---------------------------- gd 项目中的应用 ------------------------*/
func TestGDInsert(t *testing.T) {
	app.InitMongoDBClient()

	ctx := context.Background()
	logId := idutil.UniqIDV3()
	ctime := time.Now()

	var (
		orgId = int64(10)
		err   error
	)

	// 初始化dao
	dao := dao2.NewOperationLogDao(ctx, logId)

	// 插入
	insertData := &model.InsertOperationLogModel{
		OrgId:           orgId,
		UserId:          11,
		UserName:        "UserName-01",
		UserNickname:    "UserNickname-01",
		ClientIp:        "ClientIp-01",
		ReqRouter:       "ReqRouter-01",
		ReqRouterPrefix: "ReqRouterPrefix-01",
		OperateName:     "OperateName-01",
		ReqHeader:       "Sec-Fetch-Dest: empty\\nSec-Fetch-Mode: cors\\nX-Forwarded-Server: traefik-wcjdc\\nX-Forwarded-Host: apigw-test.goldentec.com\\nAccept-Language: zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2\\nReferer: https://station-test.gc365.com/\\nAccept-Encoding: gzip, deflate, br\\nOrigin: https://station-test.gc365.com\\nX-Forwarded-For: 58.251.130.162, 10.244.8.9\\nX-Forwarded-Proto: https\\nX-Forwarded-Port: 443\\nAccept: application/json, text/plain, */*\\nAccess-Token: v5_efspMtZz7p8ebAsQWIubko5fVkCFR7MC1154764563\\nContent-Type: application/json;charset=utf-8\\nContent-Length: 29\\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/118.0\\nX-Real-Ip: 10.244.8.9\\nSec-Fetch-Site: cross-site\\n",
		ReqBody:         "{\"page\":1,\"page_size\":12}",
		CostTime:        10,
		RespBody: model.RespBody{
			Code:    0,
			Message: "ok",
			LogId:   "logId-01",
			Data:    "{\"total_count\":0,\"page\":1,\"page_size\":12,\"list\":[]}",
		},
		LogId:      "logId-01",
		RespStatus: 1,
		OperatedAt: ctime.Unix(),
		CreatedAt:  ctime.Unix(),
	}
	err = dao.Insert(insertData)
	if err != nil {
		panic(err)
	}
}

func TestGDFindList(t *testing.T) {
	app.InitMongoDBClient()

	ctx := context.Background()
	logId := idutil.UniqIDV3()

	var (
		orgId = int64(10)
		err   error
	)

	// 初始化dao
	dao := dao2.NewOperationLogDao(ctx, logId)

	// 查询
	count, data, err := dao.FindList(orgId, 1, 10)
	if err != nil {
		panic(err)
	}
	dumputil.Println(count)
	dumputil.Println(data)

}

/* ---------------------------- gd 项目中的应用-end ------------------------*/
