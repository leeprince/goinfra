package cmdutil

import (
	"fmt"
	"os/exec"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/12/14 9:42
 * @Desc:
 */

// CloneGitRepo 克隆git仓库到本地路径
func CloneGitRepo(url string, dir string) error {
	cmd := exec.Command("git", "clone", url, dir)
	err := cmd.Run() // 运行命令
	if err != nil {
		return err
	}
	return nil
}

// CloneGitRepoNonBlock 克隆git仓库到本地路径|非堵塞
func CloneGitRepoNonBlock(url string, dir string) error {
	cmd := exec.Command("git", "clone", url, dir)
	err := cmd.Run() // 运行命令
	if err != nil {
		return err
	}
	
	err = cmd.Start()
	if err != nil {
		return err
	}
	
	go func() {
		err = cmd.Wait()
		if err != nil {
			fmt.Println("CloneGitRepoNoWait err:", err)
		}
		return
	}()
	
	return nil
}
