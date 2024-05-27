package base

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

var (
	TestDB *gorm.DB
)

func InitMysqlClient() {
	fmt.Println("InitMysqlClient start")

	errg := errgroup.Group{}
	errg.Go(func() error {
		// mysql.InitMysqlClient()

		return nil
	})

	err := errg.Wait()
	if err != nil {
		panic("InitMysqlClient 初始化失败" + err.Error())
	}

	fmt.Println("InitMysqlClient end")
}
