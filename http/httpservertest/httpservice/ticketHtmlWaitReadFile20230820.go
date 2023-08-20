package main

import (
	"fmt"
	"github.com/leeprince/goinfra/utils/fileutil"
	"net/http"
	"path"
	"sync"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/9 11:01
 * @Desc:
 */

// 用于解决第一次访问正常，第二次访问出现：加载页面时与服务器的连接被重置。原因每次请求都重新设置"设置静态文件目录"而导致的错误
var ticketHtmlWaitReadFile20230820Once sync.Once

func ticketHtmlWaitReadFile20230820(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	
	currentResoureDir := path.Join(resoureDir, "ticketHtmlWaitReadFile20230820")
	
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
	ticketHtmlWaitReadFile20230820Once.Do(func() {
		fs := http.FileServer(http.Dir(currentResoureDir))
		// SupplierBookingEPay.aspx.html 会定时发送请求`{当前域名}/SupplierHandler.ashx`,且不同的请求参数。
		http.Handle("/lian-tie.com_agentadmin_offline_SupplierBookingEPay.aspx_files/", http.StripPrefix("", fs))
	})
	
	fileBytes, err := fileutil.ReadFile(currentResoureDir, "lian-tie.com_agentadmin_offline_SupplierBookingEPay.aspx.html")
	if err != nil {
		http.Error(w, "读取 html文件错误", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(fileBytes))
}
