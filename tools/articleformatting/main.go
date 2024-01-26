package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()
	
	// 解决浏览器触发的同源策略错误：CORS Missing Allow Origin
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	router.Use(cors.New(config))
	
	router.POST("/format-article-file", FormatArticleFile)
	
	s := &http.Server{
		Addr:           ":18080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
