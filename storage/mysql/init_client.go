package pmysql

import (
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/6 下午10:26
 * @Desc:
 */

// 全局变量
var mysqlClients map[string]MysqlClient

type MysqlClient struct {
	db *gorm.DB
}

type MysqlConfs map[string]MysqlConf

type MysqlConf struct {
	Dsn             string           `yaml:"Dsn" json:"Dsn"`         // eg: "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	IsDebug         bool             `yaml:"IsDebug" json:"IsDebug"` // 调试启动调试模式
	MaxOpenConns    int              `yaml:"MaxOpenConns" json:"MaxOpenConns"`
	MaxIdleConns    int              `yaml:"MaxIdleConns" json:"MaxIdleConns"`
	ConnMaxLifetime time.Duration    `yaml:"ConnMaxLifetime" json:"ConnMaxLifetime"`
	Logger          logger.Interface `yaml:"-" json:"-"`
}

func InitMysqlClientMap(confs MysqlConfs) (err error) {
	mysqlClients = make(map[string]MysqlClient, len(confs))

	for name, conf := range confs {
		db, initErr := InitMysqlClient(conf)
		if initErr != nil {
			err = errors.Wrap(initErr, "InitMysqlClient error")
			return
		}

		mysqlClients[name] = MysqlClient{
			db: db,
		}
	}

	return
}

func InitMysqlClient(conf MysqlConf) (db *gorm.DB, err error) {
	if err = checkMysqlConf(conf); err != nil {
		err = errors.Wrap(err, "checkMysqlConf error")
		return
	}

	gormConfig := &gorm.Config{
		PrepareStmt: false,
		Logger:      nil,
	}
	if conf.Logger != nil {
		gormConfig.Logger = conf.Logger
	}

	db, err = gorm.Open(mysql.Open(conf.Dsn), gormConfig)
	if err != nil {
		err = errors.Wrap(err, "gorm.Open error")
		return
	}

	if conf.IsDebug {
		// 默认使用：Logger: db.Logger.LogMode(logger.Info)
		db = db.Debug()
	}

	if err = setSqlConf(db, conf); err != nil {
		err = errors.Wrap(err, "setSqlConf error")
	}

	return
}

func checkMysqlConf(conf MysqlConf) error {
	if conf.Dsn == "" {
		return errors.New("dsn must not empty")
	}
	return nil
}

func setSqlConf(db *gorm.DB, conf MysqlConf) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(MaxIdleConns)
	if conf.MaxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(conf.MaxIdleConns)
	}
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(MaxOpenConns)
	if conf.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(conf.MaxOpenConns)
	}
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(ConnMaxLifetime)
	if conf.ConnMaxLifetime > 0 {
		sqlDB.SetConnMaxLifetime(conf.ConnMaxLifetime)
	}

	return nil
}

func GetMysqlClientDB(name string) *gorm.DB {
	client, ok := mysqlClients[name]
	if !ok {
		return nil
	}
	return client.db
}
