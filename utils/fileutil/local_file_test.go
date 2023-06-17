package fileutil

import (
	"fmt"
	"testing"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/5/20 03:59
 * @Desc:
 */

func TestWriteFile(t *testing.T) {
	// dirPath := "./cache"
	dirPath := "./cache/"
	fileName := "data5.dt"
	data := []byte(fmt.Sprintf("prince-%d", time.Now().Unix()))
	isAppend := true

	gotOk, err := WriteFile(dirPath, fileName, data, isAppend)

	fmt.Println(gotOk, err)
}
