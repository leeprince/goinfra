package fileutil

import (
	"fmt"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/12/14 15:41
 * @Desc:
 */

func TestCopyFile(t *testing.T) {
	srcFile := "testcopyfile"
	dstFile := "cache/" + srcFile
	err := CopyFile(srcFile, dstFile)
	fmt.Println(err)
}
