package main

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"mime/multipart"
	"strings"
	"unicode"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/1/16 00:58
 * @Desc:
 */

func FormatArticleFile(c *gin.Context) {
	// 单文件
	file, _ := c.FormFile("file")
	if file.Size > 1000*1024*1024 {
		c.JSON(400, gin.H{
			"error": "File size cannot exceed 1000M",
		})
		return
	}
	
	headerContentType := file.Header.Get("Content-Type")
	if headerContentType != "text/plain" && headerContentType != "application/octet-stream" {
		
		log.Println("headerContentType:", headerContentType)
		c.JSON(400, gin.H{
			"error": "File type must be text",
		})
		return
	}
	
	src, err := file.Open()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to open file",
		})
		return
	}
	defer src.Close()
	
	// formattedContent, err := formatContentV1(src)
	formattedContent, err := formatContentV2(src)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	c.JSON(200, gin.H{
		"content": formattedContent,
	})
}

func formatContentV1(src multipart.File) (formattedContent string, err error) {
	content, err := io.ReadAll(src)
	if err != nil {
		return
	}
	
	formattedContent = formatContent(string(content))
	
	return
}

func formatContentV2(src multipart.File) (formattedContent string, err error) {
	scanner := bufio.NewScanner(src)
	var formattedContentBuilder strings.Builder
	for scanner.Scan() {
		line := scanner.Text()
		formattedLine := formatContent(line)
		formattedContentBuilder.WriteString(formattedLine)
		formattedContentBuilder.WriteString("\n")
	}
	
	if err = scanner.Err(); err != nil {
		return
	}
	
	formattedContent = formattedContentBuilder.String()
	
	return
}

func formatContent(content string) string {
	runes := []rune(content)
	for i := 1; i < len(runes)-1; i++ {
		if unicode.Is(unicode.Han, runes[i-1]) && unicode.IsLetter(runes[i]) {
			runes[i] = ' ' + runes[i]
		} else if unicode.IsLetter(runes[i-1]) && unicode.Is(unicode.Han, runes[i]) {
			runes[i-1] = runes[i-1] + ' '
		}
	}
	return string(runes)
}
