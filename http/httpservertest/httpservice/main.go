package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	port *int
)

var (
	// 绝对路径
	//resoureDir = "/Users/leeprince/www/go/goinfra/http/httpservertest/httpservice/resource"
	// 基于运行 main 入口函数时的相对路径; 如golang则需要看运行是指定的`Working directory`
	resoureDir = "./http/httpservertest/httpservice/resource"
)

func main() {
	port = flag.Int("port", 19999, "port")
	flag.Parse()

	// 路由定义
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/operate-html-send-request", OperateHtmlSendHttpRequest)

	// --- 项目相关
	http.HandleFunc("/ticket-html-1", ticketHtml1)
	http.HandleFunc("/ticket-html-2", ticketHtml2)

	fmt.Printf("Server listening on port:%d ...\n", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	/*
		如果在第一次访问默认路由时一切正常，但在第二次访问时浏览器报错“与服务器的连接被重置”，可能是由于服务器没有正确处理HTTP Keep-Alive 连接导致的。
		HTTP-Alive 是一种机制，它允许客户端和服务器在同一 TCP 连接上进行多个 HTTP 请求和响应。默认情况下， 的 http.Server 是启用 Keep-Alive 的，但在某些情况下，可能会出现连接被重置的问题。
		为了解决这个问题，您可以尝试在启动服务器时显式地禁用HTTP-Alive。

		我们创建了一个自定义的 http.Server 实例，并在 ConnState 回函数中禁用了 Keep-Alive。在每个新的连接状态为 http.StateNew 时，我们设置了一个较短的连接超时时间（30秒），以确保每个请求都在定时间内完成，避免连接被重置。
		请注意，这只是一种解决方案，可能不适用于所有情况。如果问题仍然存在，可能需要进一步调查和排查其他可能的原因，例如网络问题或其他服务器配置问题。
	*/
	/*server := &http.Server{
		Addr:         fmt.Sprintf(":%d", *port),
		Handler:      nil, // 使用默认的处理器
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		// 禁用 Keep-Alive
		ConnState: func(conn net.Conn, state http.ConnState) {
			if state == http.StateNew {
				conn.SetDeadline(time.Now().Add(30 * time.Second))
			}
		},
	}
	err := server.ListenAndServe()*/

	if err != nil {
		log.Fatal("ListenAndServe err:", err)
	}
}
