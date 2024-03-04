package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	
	"github.com/spf13/pflag"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/3/1 16:52
 * @Desc:
 */

func main() {
	// 定义并解析命令行标志
	titleFilePath := pflag.String("i", "", "Input path to the file containing titles")
	outputDir := pflag.String("o", ".", "Output directory for .md files")
	
	pflag.Parse()
	
	if *titleFilePath == "" {
		fmt.Println("Please provide the path to the title file using the --t flag.")
		return
	}
	
	// 打开文件
	file, err := os.Open(*titleFilePath)
	if err != nil {
		fmt.Println("Failed to open the file:", err)
		return
	}
	defer file.Close()
	
	// 创建一个扫描器来逐行读取文件
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// 获取当前行（即标题）
		title := strings.TrimSpace(scanner.Text())
		
		// 如果标题非空，则在输出目录下创建一个新的.md文件
		if title != "" {
			// 标题对应的文件名删除所有空格，但是标题保留
			titleFileName := strings.Replace(title, " ", "", -1)
			// 将特殊字符进行转义
			titleFileName = strings.Replace(titleFileName, "/", "_", -1)
			
			// 构建新的Markdown文件的完整路径
			newFilePath := filepath.Join(*outputDir, titleFileName+".md")
			
			// 创建并写入新的Markdown文件
			newFile, err := os.Create(newFilePath)
			if err != nil {
				fmt.Println("Failed to create new file:", err)
				continue
			}
			defer newFile.Close()
			
			// 在新文件中写入Markdown一级标题
			_, err = newFile.WriteString("# " + title + "\n")
			if err != nil {
				fmt.Println("Failed to write to new file:", err)
			}
		}
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	
	fmt.Println("All titles have been successfully converted to .md files in", *outputDir)
}
