package cmdutil

import (
	"fmt"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/12/14 15:46
 * @Desc:
 */

func TestCdDir(t *testing.T) {
	err := RunCmdInDir("./testtmp", "pwd")
	fmt.Println(err)
}

func TestOutputCmdInDir(t *testing.T) {
	output, err := OutputCmdInDir("./testtmp", "pwd")
	fmt.Println(string(output), err)
	
	output, err = OutputCmdInDir("./testtmp", "pwd", "-L")
	fmt.Println(string(output), err)
}
