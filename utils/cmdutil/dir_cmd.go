package cmdutil

import "os/exec"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/12/14 15:38
 * @Desc:
 */

// RunCmdInDir 在指定目录执行命令
func RunCmdInDir(dir, cmdName string, args ...string) (err error) {
	cmd := exec.Command(cmdName, args...)
	cmd.Dir = dir
	err = cmd.Run()
	return
}

// OutputCmdInDir 在指定目录执行命令并输出命令结果
func OutputCmdInDir(dir, cmdName string, args ...string) (output []byte, err error) {
	cmd := exec.Command(cmdName, args...)
	cmd.Dir = dir
	// 返回标准错误
	output, err = cmd.Output()
	// 返回其组合的标准输出和标准错误：调试时很好用
	// output, err = cmd.CombinedOutput()
	if err != nil {
		return
	}
	return
}
