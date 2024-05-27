package base

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/5/27 17:13
 * @Desc:
 */

func Init(env string) (err error) {
	err = InitConfig(env)
	if err != nil {
		return
	}

	InitLog()

	InitMysqlClient()
	if err != nil {
		return
	}
	return
}
