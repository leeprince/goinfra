package signalutil

import (
	"log"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/9/2 13:31
 * @Desc:
 */

func TestInitOsSigal(t *testing.T) {
	InitOsSigal()
	log.Println("InitOsSigal success!")
	
	for {
		select {
		case sigal, ok := <-OsSigal:
			if !ok {
				log.Println("<-chan os.Signal !ok")
			}
			log.Println("<-chan os.Signal:", sigal)
			return
		}
	}
}

func TestInitOsSigalMore(t *testing.T) {
	InitOsSigal()
	log.Println("TestInitOsSigalMore success!")
	
	// 多个监听 OsSigal： 只会有一个接收到信道。并且该场景下，只有sigal2能接收到，跟GMP有关
	go func() {
		for {
			select {
			case sigal, ok := <-OsSigal:
				if !ok {
					log.Println("<-chan os.Signal !ok")
				}
				log.Println("<-chan os.Signal:", sigal)
				return
			}
		}
	}()
	
	for {
		select {
		case sigal, ok := <-OsSigal:
			if !ok {
				log.Println("<-chan os.Signal !ok")
			}
			log.Println("<-chan os.Signal:", sigal)
			return
		}
	}
}
