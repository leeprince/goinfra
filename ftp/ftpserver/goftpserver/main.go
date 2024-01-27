package main

import (
	filedriver "github.com/goftp/file-driver"
	"github.com/goftp/server"
	"log"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/1/27 10:18
 * @Desc:
 */

const (
	RootPath     = "/Users/leeprince/www/go/goinfra/ftp/ftpserver/goftpserver/tmp"
	Port         = 2121
	AuthUsername = "admin"
	AuthPassword = "123456"
)

func main() {
	factory := &filedriver.FileDriverFactory{
		RootPath: RootPath,
		Perm:     server.NewSimplePerm("user", "group"),
	}
	
	opts := &server.ServerOpts{
		Factory:  factory,
		Port:     Port,
		Hostname: "localhost",
		Auth:     &server.SimpleAuth{Name: AuthUsername, Password: AuthPassword},
	}
	
	ftpServer := server.NewServer(opts)
	err := ftpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
