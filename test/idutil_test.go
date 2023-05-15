package test

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

func TestUUID(t *testing.T) {
	var id string
	id = idutil.UUID()
	fmt.Println("UUID:", id)
}

func TestSnowflakeID(t *testing.T) {
	var idNO int64
	idNO = idutil.NewSnowflake(10).NextId()
	fmt.Println("NewSnowflake NextId int64:", idNO)
	id := fmt.Sprintf("%016x", idNO)
	fmt.Println("NewSnowflake NextId string:", id)
}
func BenchmarkSnowflakeID(b *testing.B) {
	var idNO int64
	for i := 0; i < b.N; i++ {
		idNO = idutil.NewSnowflake(1).NextId()
		fmt.Sprintf("%016x", idNO)
	}
}

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
