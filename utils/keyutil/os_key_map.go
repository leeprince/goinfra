package keyutil

import (
	"errors"
	"github.com/leeprince/goinfra/consts"
	"github.com/leeprince/goinfra/consts/constval"
	"github.com/leeprince/goinfra/utils/arrayutil"
	"runtime"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/21 18:40
 * @Desc:
 */

func GetOsKeyButtonRawCode() (keyMap *constval.StringUint16Group, err error) {
	if !arrayutil.InString(runtime.GOOS, []string{
		consts.GOOSDarwin,
		consts.GOOSWindows,
	}) {
		err = errors.New("暂不支持该操作系统")
		return
	}
	keyMap = consts.WindowsOSKeyButtonRawcode
	if runtime.GOOS == consts.GOOSDarwin {
		keyMap = consts.DarwinOSKeyButtonRawcode
	}
	return
}
