package fileutil

import (
	"io"
	"os"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/12/14 15:14
 * @Desc:
 */

func CopyFile(srcFile string, dstFile string) error {
	// 打开源文件
	src, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer src.Close()
	
	// 创建目标文件
	dst, err := os.Create(dstFile)
	if err != nil {
		return err
	}
	defer dst.Close()
	
	// 复制文件
	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}
	
	// 刷新文件缓存
	err = dst.Sync()
	if err != nil {
		return err
	}
	
	return nil
}
