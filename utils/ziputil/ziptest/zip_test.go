package ziptest

import (
	"archive/zip"
	"io"
	"log"
	"sync"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/1 16:13
 * @Desc:
 */

func TestReadZIP(t *testing.T) {
	pathFile := "./dzfp_23612000000006098062_20230517164634.zip"

	// 解压zip文件
	zipReader, err := zip.OpenReader(pathFile)
	if err != nil {
		log.Fatal(err)
	}
	defer zipReader.Close()

	wg := sync.WaitGroup{}
	for _, file := range zipReader.File {
		fileItem := file
		if fileItem.FileInfo().IsDir() {
			log.Println("IsDir")
			return
		}

		wg.Add(1)
		go func(fileItem *zip.File) {
			defer wg.Done()

			fileNameRaw := fileItem.FileInfo().Name()
			log.Println(fileNameRaw)

			//fileRaw, err := fileItem.OpenRaw() // 读取到乱码数据
			fileRaw, err := fileItem.Open()
			if err != nil {
				log.Fatal("fileItem.Open()")
				return
			}
			defer func(fileRaw io.ReadCloser) {
				fileRaw.Close()
			}(fileRaw)

			fileRawBytes, err := io.ReadAll(fileRaw)
			if err != nil {
				log.Fatal("fileItem.Open()")
				return
			}

			log.Println("fileNameRaw:", fileNameRaw)
			log.Println("fileRaw:", fileRaw)
			log.Println("fileRawBytes:", string(fileRawBytes))
		}(fileItem)
	}
	wg.Wait()
}
