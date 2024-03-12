package main

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/3/9 03:12
 * @Desc:
 */
import (
	"fmt"
	"github.com/spf13/pflag"
	"os"
	"path/filepath"
	"strings"
)

// 遍历指定目录下的所有文件，并将文件名中的 '?' 替换为 '-'：
func main() {
	// 指定你的目录路径
	inputDir := pflag.String("dir", "", "指定你的目录路径")
	pflag.Parse()
	
	dir := *inputDir
	if dir == "" {
		fmt.Println("请指定目录路径")
		return
	}
	return
	
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			oldFilename := info.Name()
			newFilename := strings.Replace(oldFilename, "?", "-", -1)
			if oldFilename != newFilename {
				oldPath := filepath.Join(dir, oldFilename)
				newPath := filepath.Join(dir, newFilename)
				err := os.Rename(oldPath, newPath)
				if err != nil {
					fmt.Printf("Failed to rename %s to %s: %v\n", oldPath, newPath, err)
					return err
				}
				fmt.Printf("Successfully renamed %s to %s\n", oldPath, newPath)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error walking the path:", err)
	}
}
