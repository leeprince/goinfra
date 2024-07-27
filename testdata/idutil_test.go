package testdata

import (
	"fmt"
	"github.com/leeprince/goinfra/utils/idutil"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/16 上午10:34
 * @Desc:
 */

func TestUniqID(t *testing.T) {
	fmt.Println(idutil.UniqIDV1())
	fmt.Println(idutil.UniqIDV2())
	fmt.Println(idutil.UniqIDV3())
}

func BenchmarkUniqIDV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		idutil.UniqIDV1()
	}
}

func BenchmarkUniqIDV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		idutil.UniqIDV2()
	}
}

func BenchmarkUniqIDV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		idutil.UniqIDV3()
	}
}

func BenchmarkUniqIDV4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		idutil.UniqIDV4()
	}
}
