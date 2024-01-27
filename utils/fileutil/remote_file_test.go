package fileutil

import (
	"fmt"
	"sync"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/1/26 18:33
 * @Desc:
 */

func TestWriteFileByUrl(t *testing.T) {
	urls := []string{}
	
	exitUrl := make(map[string]bool)
	for _, url := range urls {
		if exitUrl[url] {
			continue
		}
		exitUrl[url] = true
	}
	
	var wg sync.WaitGroup
	for url, _ := range exitUrl {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			filename := GetFileNameFromURL(url)
			filePath := "./tmp/"
			gotPathFile, err := WriteFileByUrl(url, filename, filePath)
			if err != nil {
				fmt.Println("WriteFileByUrl err", err)
			}
			fmt.Println("WriteFileByUrl gotPathFile", gotPathFile)
			
		}(url)
	}
	wg.Wait()
	
	fmt.Println("--- TestWriteFileByUrl")
}
