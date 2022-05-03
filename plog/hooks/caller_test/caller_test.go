package caller_test

import (
    "github.com/leeprince/goinfra/plog/hooks/caller_test/callerfile"
    "testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/15 上午10:19
 * @Desc:
 */

func TestCaller(t *testing.T)  {
    for i := 0; i <= 5; i++ {
        callerfile.Caller(i)
    }
}