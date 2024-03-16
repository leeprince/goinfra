package api

import "getwebpage-tomarkdown/internel/application"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/3/16 02:33
 * @Desc:
 */

type API interface {
	SaveWebPage()
}

func NewAPI() {
	// 初始化应用层
	app := application.NewApplication()
	app.SaveWebPage()
}
