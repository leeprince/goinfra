package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/leeprince/goinfra/config"
	"github.com/sogko/go-wordpress"
	"github.com/spf13/pflag"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/3/3 01:59
 * @Desc:
 */

var (
	titleFileDir *string
	catID        *int
	tagID        *int
)

func init() {
	config.InitConfig("./config/config.yaml")
	
	// 定义并解析命令行标志：帮助命令：go run main.go --help
	titleFileDir = pflag.String("dir", "./tools/wordpress/gowordpress/resource", "输入包含需要提交文件的目录路径")
	catID = pflag.Int("catid", 152, "该批次所属分类ID")
	tagID = pflag.Int("tagid", 153, "该批次所属标签ID")
	
	pflag.Parse()
	
	if *titleFileDir == "" {
		fmt.Println("Please provide the path to the title file using the --i flag.")
		return
	}
}

func main() {
	// 验证 SDK 有效性：获取用户列表
	// verifySdkValid()
	
	// 列出文章
	// ListPost()
	
	// 创建文章
	// CreatePost("皇子谈技术", "# 概述\n# 架构设计 # 拥抱 AI", []int{152}, []int{153})
	
	// 批量创建文章
	batchCreatePost()
}

//
func batchCreatePost() {
	var errArr []error
	err := filepath.WalkDir(*titleFileDir, func(path string, info fs.DirEntry, err error) (newError error) {
		if err != nil {
			fmt.Println("Error walking path:", err)
			err = errors.New(fmt.Sprintf("path: %s Walk error: %s", path, err.Error()))
			errArr = append(errArr, err)
			return
		}
		
		if info.IsDir() {
			fmt.Println("info IsDir:", info.Name())
			return
		}
		
		if !strings.HasSuffix(path, ".md") {
			fmt.Println("Markdown file not found:", path)
			return
		}
		
		// 打开文件
		file, err := os.Open(path)
		if err != nil {
			fmt.Println("Failed to open the file:", err)
			err = errors.New(fmt.Sprintf("path: %s Open error: %s", path, err.Error()))
			errArr = append(errArr, err)
			return
		}
		defer file.Close()
		
		// 获取标题
		scanner := bufio.NewScanner(file)
		contentBuilder := strings.Builder{}
		var title string
		for scanner.Scan() {
			// 获取标题
			if title == "" {
				// 获取第一行为空，即标题
				title = strings.TrimSpace(scanner.Text())
				
				if title == "" {
					continue
				}
				title = strings.ReplaceAll(title, "# ", "")
			}
			
			// 写入内容
			lineContent := strings.ReplaceAll(scanner.Text(), "(typora/", "(/typora/")
			contentBuilder.WriteString(lineContent + "\n")
		}
		if err = scanner.Err(); err != nil {
			err = errors.New(fmt.Sprintf("path: %s scanner error: %s", path, err.Error()))
			errArr = append(errArr, err)
			return
		}
		content := contentBuilder.String()
		
		err = CreatePost(title, content, []int{*catID}, []int{*tagID})
		if err != nil {
			err = errors.New(fmt.Sprintf("path: %s CreatePost error: %s", path, err.Error()))
			errArr = append(errArr, err)
			return
		}
		
		return
	})
	
	if err != nil {
		fmt.Println("Error walking path:", err)
	}
}

func CreatePost(titleRaw, contentRaw string, catList, tagList []int) (err error) {
	fmt.Println("titleRaw:", titleRaw, "catList:", catList, "tagList:", tagList)
	
	// 创建帖子（文章）：https://developer.wordpress.org/rest-api/reference/posts/#create-a-post
	getParamsOrPostContent := struct {
		wordpress.Post
		Categories []int `json:"categories,omitempty"` // 已存在的分类ID
		Tags       []int `json:"tags,omitempty"`       // 已存在的标签ID
	}{
		Post: wordpress.Post{
			ID:          0,
			Date:        "",
			DateGMT:     "",
			GUID:        wordpress.GUID{},
			Link:        "",
			Modified:    "",
			ModifiedGMT: "",
			Password:    "",
			Slug:        "",
			Status:      wordpress.PostStatusPublish,
			Type:        wordpress.PostTypePost,
			Title: wordpress.Title{
				Raw:      titleRaw,
				Rendered: "",
			},
			Content: wordpress.Content{
				Raw:      contentRaw,
				Rendered: "",
			},
			Author:        0,
			Excerpt:       wordpress.Excerpt{},
			FeaturedImage: 0,
			CommentStatus: "",
			PingStatus:    "",
			Format:        wordpress.PostFormatStandard,
			Sticky:        false,
		},
		Categories: catList,
		Tags:       tagList,
	}
	
	// sdk 没兼容好 url，暂未能使用
	/*client := getSdkClient()
	newPost, resp, body, err := client.Posts().Create(&getParamsOrPostContent)
	if err != nil {
		fmt.Println("client request err:", err)
		return
	}
	fmt.Printf("- resp：%+v \n", resp)
	if resp.StatusCode != http.StatusOK {
		fmt.Println("client request StatusCode not StatusOK")
		return
	}
	fmt.Println("- body：", string(body))
	fmt.Println("- newPost：", newPost)
	return*/
	
	url := "/wp/v2/posts"
	var result interface{}
	_, _, err = sdkCenter(url, "create", getParamsOrPostContent, &result)
	if err != nil {
		fmt.Printf("sdkCenter title:%s, err: %v", titleRaw, err)
		return
	}
	return
}

// 获取文章列表
func ListPost() {
	// 帖子（文章）列表：https://developer.wordpress.org/rest-api/reference/posts/#list-posts
	url := "/wp/v2/posts"
	var result interface{}
	_, _, err := sdkCenter(url, "get", nil, &result)
	if err != nil {
		log.Fatal(err)
	}
}

// 验证 SDK 有效性：获取用户列表
func verifySdkValid() {
	// 获取用户列表：
	url := "/wp/v2/users"
	var result interface{}
	_, _, err := sdkCenter(url, "get", nil, &result)
	if err != nil {
		log.Fatal(err)
	}
}

// method: get/create
func sdkCenter(path, method string, getParamsOrPostContent, result interface{}) (resp *http.Response, body []byte, err error) {
	fmt.Printf("path: %s；method:%s \n", path, method)
	if path == "" {
		err = errors.New("path is nil")
		return
	}
	if method == "" {
		err = errors.New("method is nil")
		return
	}
	if result == nil {
		err = errors.New("result is nil")
		return
	}
	
	apiPrefix := "/wp-json"
	url := fmt.Sprintf("%s%s%s", strings.TrimRight(config.C.WordPress.Host, "/"), apiPrefix, path)
	fmt.Printf("full url: %s \n", url)
	
	// 创建客户端
	client := getSdkClient()
	
	method = strings.ToUpper(method)
	switch method {
	case "GET":
		resp, body, err = client.Get(url, getParamsOrPostContent, result)
	case "CREATE":
		resp, body, err = client.Create(url, getParamsOrPostContent, result)
	default:
		err = errors.New("method not support")
		return
	}
	if err != nil {
		fmt.Println("client request err:", err)
		return
	}
	fmt.Printf("- resp：%+v \n", resp)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		fmt.Println("client request StatusCode not StatusOK")
		return
	}
	fmt.Println("- body：", string(body))
	return
}

// 创建客户端
func getSdkClient() *wordpress.Client {
	// 创建客户端
	// 使用用户和密码的方式进行验证：https://developer.wordpress.org/rest-api/using-the-rest-api/authentication/
	client := wordpress.NewClient(&wordpress.Options{
		BaseAPIURL: config.C.WordPress.Host,
		Username:   config.C.WordPress.Username,
		Password:   config.C.WordPress.Password,
	})
	return client
}
