package rartest

import (
	"github.com/mholt/archiver/v3"
	"log"
	"os"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/1 14:32
 * @Desc:
 */
func TestReadRAR(t *testing.T) {
	// 解压RAR文件
	err := archiver.Unarchive("dzfp_23612000000006098062_20230517164634.rar", "./tmp/")
	if err != nil {
		log.Fatal(err)
	}

	// 读取RAR文件中的文件
	files, err := os.ReadDir("./tmp/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			data, err := os.ReadFile("./tmp/" + file.Name())
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("file: %s, content: %s\n", file.Name(), string(data))
		}
	}
}
