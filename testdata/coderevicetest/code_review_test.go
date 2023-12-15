package coderevicetest

import (
	"fmt"
	"github.com/leeprince/goinfra/utils/cmdutil"
	"github.com/leeprince/goinfra/utils/fileutil"
	"log"
	"path/filepath"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/12/14 10:38
 * @Desc:
 */

var gitRepoUrlLocalPathMap = map[string]string{
	"https://github.com/leeprince/go_micro.git": "./tmp/go_micro",
}

func TestCodeReview(t *testing.T) {
	for repoUrl, localPath := range gitRepoUrlLocalPathMap {
		// 克隆项目到本地临时文件夹
		repoUrl := repoUrl
		localPath := localPath
		// 配置时路径包含仓库名
		/*repoName, err := urlutil.GetRepoName(repoUrl)
		if err != nil {
			log.Fatal("urlutil.GetRepoName err:", err)
		}
		localPath = filepath.Join(localPath, repoName)*/
		fmt.Println("localPath:", localPath)
		// 检查目录是否已存在
		_, ok := fileutil.CheckFileDirExist(localPath)
		if ok {
			// 检查master>main>test分支是否存在
			cmdutil.OutputCmdInDir(localPath, "git", "branch")
			
			// 按优先级master>main>test拉取分支
			
			// 判断是否有代码提交：无代码提交则终止审查；有代码提交则继续审查
			
		} else {
			// 克隆项目到本地
			err := cmdutil.CloneGitRepo(repoUrl, localPath)
			if err != nil {
				log.Fatal("gitutil.CloneGitRepo err:", err)
			}
			// 复制代码审查文件到项目中
			srcFile := "Makefile"
			dstFile := filepath.Join(localPath, "Makefile")
			err = fileutil.CopyFile(srcFile, dstFile)
			if err != nil {
				log.Fatal("fileutil.CopyFile err:", err)
			}
			
		}
		
		// 执行代码审查
		output, err := cmdutil.OutputCmdInDir(localPath, "make", "vet")
		if err != nil {
			log.Fatal("cmdutil.OutputCmdInDir(make vet err:", err)
		}
		fmt.Println("output:", output)
	}
}

func TestConcurrentCodeReview(t *testing.T) {
	// var errorArr []error
	// var mtLock sync.Mutex
	for repoUrl, localPath := range gitRepoUrlLocalPathMap {
		repoUrl := repoUrl
		localPath := localPath
		/*go func() {
			err := gitutil.CloneGitRepo(repoUrl, localPath)
			if err != nil {
				ok := mtLock.TryLock()
				if !ok {
					log.Fatal("mtLock.TryLock() !ok")
					return
				}
				defer mtLock.Unlock()
				errorArr = append(errorArr, err)
			}
		}()*/
		err := cmdutil.CloneGitRepo(repoUrl, localPath)
		if err != nil {
			log.Fatal("gitutil.CloneGitRepo err:", err)
		}
	}
}

func TestName(t *testing.T) {
	localPath := "./tmp"
	repoName := "hello"
	localPath = filepath.Join(localPath, repoName)
	fmt.Println(localPath)
}

func TestCopy(t *testing.T) {
	err := fileutil.CopyFile("Makefile", "tmp/go_micro/Makefile")
	if err != nil {
		log.Fatal("fileutil.CopyFile err:", err)
	}
}
