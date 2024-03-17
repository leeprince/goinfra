package fileutil

import (
	"fmt"
	"path/filepath"
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

func TestFilePath(t *testing.T) {
	a1 := "a1"
	a2 := "/a2"
	a3 := "a3/"
	a4 := "/a4/"
	
	a := filepath.Join(a1, a2, a3, a4)
	fmt.Println(a)
}
