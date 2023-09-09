package envutil

import (
	"fmt"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/9 17:49
 * @Desc:
 */

func TestGetenvIsMock(t *testing.T) {
	isMock := EnvIsMock()
	fmt.Println("isMock:", isMock)
}
