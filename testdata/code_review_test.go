package testdata

import (
	"fmt"
	"github.com/leeprince/goinfra/utils/gitutil"
	"github.com/leeprince/goinfra/utils/urlutil"
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
	"https://github.com/leeprince/go_micro.git": "./tmp",
}

func TestCodeReview(t *testing.T) {
	for repoUrl, localPath := range gitRepoUrlLocalPathMap {
		repoUrl := repoUrl
		localPath := localPath
		
		repoName, err := urlutil.GetRepoName(repoUrl)
		if err != nil {
			log.Fatal("urlutil.GetRepoName err:", err)
		}
		
		localPath = filepath.Join(localPath, repoName)
		fmt.Println("localPath:", localPath)
		err = gitutil.CloneRepo(repoUrl, localPath)
		if err != nil {
			log.Fatal("gitutil.CloneRepo err:", err)
		}
	}
}

func TestConcurrentCodeReview(t *testing.T) {
	// var errorArr []error
	// var mtLock sync.Mutex
	for repoUrl, localPath := range gitRepoUrlLocalPathMap {
		repoUrl := repoUrl
		localPath := localPath
		/*go func() {
			err := gitutil.CloneRepo(repoUrl, localPath)
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
		err := gitutil.CloneRepo(repoUrl, localPath)
		if err != nil {
			log.Fatal("gitutil.CloneRepo err:", err)
		}
	}
}

func TestName(t *testing.T) {
	localPath := "./tmp"
	repoName := "hello"
	localPath = filepath.Join(localPath, repoName)
	fmt.Println(localPath)
}
