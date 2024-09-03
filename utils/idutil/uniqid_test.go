package idutil

import (
	"fmt"
	"sync"
	"testing"
	"time"
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

func TestGenerate(t *testing.T) {
	var wg sync.WaitGroup
	snowflakeGen := &SnowflakeGenerator{}
	startTime := time.Now()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			id := snowflakeGen.Generate()
			fmt.Printf("Generated ID: %d\n", id)
		}()
	}
	wg.Wait()
	endTime := time.Now()
	fmt.Printf("Generated 100 IDs in %v\n", endTime.Sub(startTime))
}
