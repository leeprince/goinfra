package gitutil

import "os/exec"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/12/14 9:42
 * @Desc:
 */

// CloneRepo 克隆git仓库到本地路径
func CloneRepo(url string, dir string) error {
	cmd := exec.Command("git", "clone", url, dir)
	err := cmd.Run() // 运行命令
	if err != nil {
		return err
	}
	return nil
}
