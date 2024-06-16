package idutil

import (
	"fmt"
	"sync"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/16 上午10:34
 * @Desc:
 */

func TestUUID(t *testing.T) {
	var id string
	id = UniqIDV2()
	fmt.Println("UUID:", id)
}

func TestMaxWorker(t *testing.T) {
	fmt.Println(maxWorker)
}
func TestSnowflakeID(t *testing.T) {
	var idNO int64
	idNO = NewSnowflake(10).NextId()
	fmt.Println("NewSnowflake NextId int64:", idNO)
	id := fmt.Sprintf("%016x", idNO)
	fmt.Println("NewSnowflake NextId string:", id)
}
func BenchmarkSnowflakeID(b *testing.B) {
	var idNO int64
	for i := 0; i < b.N; i++ {
		idNO = NewSnowflake(1).NextId()
		fmt.Sprintf("%016x", idNO)
	}
}

func TestUniqID(t *testing.T) {
	fmt.Println(UniqIDV1())
	fmt.Println(UniqIDV2())
	fmt.Println(UniqIDV3())
	
	fmt.Println()
	// time.Sleep(time.Millisecond * 1)
	
	fmt.Println(UniqIDV1())
	fmt.Println(UniqIDV2())
	fmt.Println(UniqIDV3())
	
	fmt.Println()
	// time.Sleep(time.Millisecond * 1)
	
	fmt.Println(UniqIDV1())
	fmt.Println(UniqIDV2())
	fmt.Println(UniqIDV3())
}

func BenchmarkUniqIDV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UniqIDV1()
	}
}

func BenchmarkUniqIDV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UniqIDV2()
	}
}

func BenchmarkUniqIDV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UniqIDV3()
	}
}

func TestUniqIDV3Goroutine(T *testing.T) {
	count := 100
	var wg sync.WaitGroup
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			fmt.Println(UniqIDV3())
		}()
	}
	wg.Wait()
}
