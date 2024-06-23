package pmysql

import (
	"context"
	"fmt"
	"github.com/leeprince/goinfra/plog"
	"github.com/leeprince/goinfra/storage/mysql/gorm_test_model"
	"gorm.io/gorm/logger"
	"testing"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/11 下午9:52
 * @Desc:
 */

func initLoger() {
	// err := plog.SetOutputFile("./logs", "gorm.log", false)
	err := plog.SetOutputFile("./gorm_test_log", "gorm.log", true)
	if err != nil {
		panic(fmt.Sprintf("plog.SetOutputFile err:%v", err))
	}
	// plog.SetReportCaller(true)
}

func TestInitMysqlClient(t *testing.T) {
	initLoger()
	// var logWriterStdout = log.New(os.Stdout, "\r\n", log.LstdFlags) // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
	var DBLogger = logger.New(
		// logWriterStdout, // 标准输出
		plog.GetLogger(), // 指定日志文件输出
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Warn, // 日志级别
			IgnoreRecordNotFoundError: false,       // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 彩色打印
		},
	)
	
	tmpMysql := "tmp"
	
	type args struct {
		confs MysqlConfMap
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				confs: MysqlConfMap{
					tmpMysql: MysqlConf{
						Dsn:     "root:leeprince@tcp(127.0.0.1:3306)/tmp?charset=utf8mb4&parseTime=true&loc=Local&interpolateParams=True",
						IsDebug: true,
						Logger:  DBLogger,
					},
				}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitMysqlClientMap(tt.args.confs); (err != nil) != tt.wantErr {
				t.Errorf("InitMysqlClientMap() error = %v, wantErr %v", err, tt.wantErr)
			}
			
			tmpDB := GetMysqlClientDB(tmpMysql)
			if tmpDB == nil {
				fmt.Println(" GetMysqlClientDB(tmpMysql)")
				return
			}
			
			userDao := gorm_test_model.NewUsersDao(context.Background(), tmpDB)
			user, err := userDao.GetByOption(userDao.WithID(1))
			fmt.Println("GetByOption:", user, err)
			findName := "name01"
			users, err := userDao.GetByOptions(userDao.WithName(&findName))
			fmt.Println("GetByOptions:", users, err)
			
		})
	}
}
