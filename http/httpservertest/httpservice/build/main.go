package main

import (
	"flag"
	"fmt"
	"github.com/leeprince/goinfra/utils/fileutil"
	"log"
	"net/http"
	"path"
	"sync"
)

var (
	port *int
)

var (
	resoureDir = "./http/httpservertest/httpservice/resource"
)

func main() {
	port = flag.Int("port", 19999, "port")
	flag.Parse()
	
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

// 用于解决第一次访问正常，第二次访问出现：加载页面时与服务器的连接被重置。原因每次请求都重新设置"设置静态文件目录"而导致的错误
var ticketHtml2Once sync.Once

func ticketHtml2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	
	resoureDir := path.Join(resoureDir, "html8")
	
	// 设置静态文件目录
	/*
		http.FileServer函数创建了一个处理静态文件的处理器，并将其注册到以/test/为前缀的路由上。这样，当访问以/test/开头的URL时，服务器将在定的静态文件目录中查找相应的文件并返回给客户端。
		通过设置静态文件目录，您可以将静态资源（如CSS、JavaScript、图像文件等）与您的服务器代码分开存放。这样做的好处是：
		1. 简化代码：将静态文件的处理交给专门的处理器，使得您的服务器代码更加简洁和专注于业务逻辑。
		2. 提高性能：静态文件可以由Web服务器直接提供，而不需要经过您的应用程序处理。这样可以减轻应用程序的负载，提高响应速度和并发能力。
		3. 组织结构清晰：将静态文件放在单独的目录中，使得项目结构更加清晰，易于维护和扩展。
	
		例如：
		例如，如果有一个名为test/style.css的CSS文件，可以通过访问http://localhost:8080/test/style.css来获取该文件的内容。
		通过设置静态文件目录，您可以轻松地为您的Web应用程序提供CSS、JavaScript、图像等静态，并且可以根据需要进行组织和管理。
		在http.Handle("/test/", http.StripPrefix("/test/", fs))中，第一个"/test/"表示要处理的URL前缀，第二个"/test/"表示要从URL剥离的前缀。具体来说，这个语句的作用是将以/test/开的URL请求转发到fs处理器，并从URL中剥离掉/test/前缀。
		http.StripPrefix函数的作用是从URL中剥离指定的前缀。在这个例子中，我们使用http.StripPrefix("/test/", fs)将fs处理器注册到以/test/开头的URL上，并从URL中剥离掉/test/前缀。这样，当客户端请求/test/style.css时，服务器会将请求转发给fs处理器，并从URL中剥离掉/test/前缀，最终静态文件目录中查找名为style.css的文件并返回给客户端。
		总之，http.Handle("/test/", http.StripPrefix("/test/", fs))的作用是将以/test/开头的URL请求转发到fs处理器，并从URL中剥离掉/test/前缀，以便在静态文件目录中查找相应的文件。
	*/
	ticketHtml2Once.Do(func() {
		fs := http.FileServer(http.Dir(resoureDir))
		http.Handle("/SupplierBookingEPay.aspx_files/", http.StripPrefix("", fs))
	})
	
	fileBytes, err := fileutil.ReadFile(resoureDir, "SupplierBookingEPay.aspx.html")
	if err != nil {
		http.Error(w, "读取 html文件错误", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(fileBytes))
}
