package urlutil

import (
	"net/url"
	"path"
	"strings"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/12/14 11:07
 * @Desc:
 */

// 仅根据一个github url获取到仓库名。github url格式为：https://github.com/leeprince/go_micro或者 https://github.com/leeprince/go_micro.git
func GetRepoName(repoURL string) (string, error) {
	u, err := url.Parse(repoURL)
	if err != nil {
		return "", err
	}
	
	repoNameWithExt := path.Base(u.Path)                    // 获取路径的最后一部分，可能包含 ".git"
	repoName := strings.TrimSuffix(repoNameWithExt, ".git") // 移除可能的 ".git" 后缀
	
	return repoName, nil
}
