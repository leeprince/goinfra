package ppostgres

import (
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
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
var mysqlClients map[string]PostgresClient

type PostgresClient struct {
	db *gorm.DB
}

type PostgresConfMap map[string]PostgresConf

type PostgresConf struct {
	Dsn             string           `yaml:"Dsn" json:"Dsn"`                         // eg: "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	IsDebug         bool             `yaml:"IsDebug" json:"IsDebug"`                 // 调试启动调试模式
	MaxOpenConns    int              `yaml:"MaxOpenConns" json:"MaxOpenConns"`       // 设置打开数据库连接的最大数量
	MaxIdleConns    int              `yaml:"MaxIdleConns" json:"MaxIdleConns"`       // 设置空闲连接池中连接的最大数量
	ConnMaxLifetime time.Duration    `yaml:"ConnMaxLifetime" json:"ConnMaxLifetime"` // 设置了连接可复用的最大时间
	Logger          logger.Interface `yaml:"-" json:"-"`                             // SQL 日志
}

func InitPostgresClientMap(confs PostgresConfMap) (err error) {
	mysqlClients = make(map[string]PostgresClient, len(confs))
	
	for name, conf := range confs {
		db, initErr := InitPostgresClient(conf)
		if initErr != nil {
			err = errors.Wrap(initErr, "InitPostgresClient error")
			return
		}
		
		mysqlClients[name] = PostgresClient{
			db: db,
		}
	}
	
	return
}

func MustInitPostgresClient(conf PostgresConf) (db *gorm.DB) {
	db, err := InitPostgresClient(conf)
	if err != nil {
		panic("初始化数据库连接池失败：" + err.Error())
	}
	return db
}

func InitPostgresClient(conf PostgresConf) (db *gorm.DB, err error) {
	if err = checkPostgresConf(conf); err != nil {
		err = errors.Wrap(err, "checkPostgresConf error")
		return
	}
	
	gormConfig := &gorm.Config{
		PrepareStmt: false,
		Logger:      nil,
	}
	if conf.Logger != nil {
		gormConfig.Logger = conf.Logger
	}
	
	db, err = gorm.Open(postgres.Open(conf.Dsn), gormConfig)
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

func checkPostgresConf(conf PostgresConf) error {
	if conf.Dsn == "" {
		return errors.New("dsn must not empty")
	}
	return nil
}

func setSqlConf(db *gorm.DB, conf PostgresConf) error {
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
