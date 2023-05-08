package timeutil

import (
	"fmt"
	"github.com/leeprince/goinfra/consts"
	"sync"
	"testing"
	"time"
)

/**
 * @Author: prince.lee
 * @Date:   2022/3/24 17:13
 * @Desc:
 */

func TestUseMillisecondUnit(t *testing.T) {
	type args struct {
		dur time.Duration
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				dur: 0,
			},
			want: true,
		},
		{
			args: args{
				dur: -1,
			},
			want: true,
		},
		{
			args: args{
				dur: time.Millisecond * 100,
			},
			want: true,
		},
		{
			args: args{
				dur: time.Millisecond * 1000,
			},
		},
		{
			args: args{
				dur: time.Millisecond * 2000,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UseMillisecondUnit(tt.args.dur); got != tt.want {
				t.Errorf("UseMillisecondUnit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeTick(t *testing.T) {
	const goroutines = 1
	var wg sync.WaitGroup
	wg.Add(goroutines)
	var mu sync.Mutex
	ticker := time.Tick(time.Second * 3)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for c := 0; c < 1000; c++ {
				mu.Lock()
				time.Sleep(time.Second * 1)
				fmt.Println("cccccc:", c)
				mu.Unlock()
			}

		}()
	}

	// 读出通道方式一【推荐】
	for {
		if v, ok := <-ticker; ok {
			fmt.Println(">>>>>>>>", v, ok)
			// return // 测试time.Tick
		}
	}
	// 读出通道方式二 // 此处只有一个通道等待，推荐使用方式一
	// for {
	//     select {
	//     case <-ticker:
	//         fmt.Println(">>>>>>>>")
	//         // return // 测试time.Tick
	//     }
	// }

	wg.Wait()
}

func TestNewTicker(t *testing.T) {
	const goroutines = 1
	var wg sync.WaitGroup
	wg.Add(goroutines)
	var mu sync.Mutex
	ticker := time.NewTicker(time.Second * 3)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for c := 0; c < 1000; c++ {
				mu.Lock()
				time.Sleep(time.Second * 1)
				fmt.Println("cccccc:", c)
				mu.Unlock()
			}

		}()
	}

	// 读出通道方式一【推荐】
	for {
		if v, ok := <-ticker.C; ok {
			fmt.Println(">>>>>>>>", v, ok)

			ticker.Stop() // 测试ticker.Stop()
			// ticker.Reset(time.Second * 1) // 测试ticker.Stop()

			// return // 测试time.NewTicker
		}
	}
	// 读出通道方式二 // 此处只有一个通道等待，推荐使用方式一
	// for {
	//     select {
	//     case <-ticker:
	//         fmt.Println(">>>>>>>>")
	//         // return // 测试time.Tick
	//     }
	// }

	wg.Wait()
}

func TestTimeAfter(t *testing.T) {
	const goroutines = 1
	var wg sync.WaitGroup
	wg.Add(goroutines)
	var mu sync.Mutex
	ticker := time.After(time.Second * 3) // time.After == time.NewTimer.C

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for c := 0; c < 1000; c++ {
				mu.Lock()
				time.Sleep(time.Second * 1)
				fmt.Println("cccccc:", c)
				mu.Unlock()
			}
		}()
	}

	// 读出通道方式一【推荐】
	for {
		if v, ok := <-ticker; ok {
			fmt.Println(">>>>>>>>", v, ok)
			// return // 测试ticker
		}
	}
	// 读出通道方式二 // 此处只有一个通道等待，推荐使用方式一
	// for {
	//     select {
	//     case <-ticker:
	//         fmt.Println(">>>>>>>>")
	//         // return // 测试ticker
	//     }
	// }

	wg.Wait()
}

func TestNewTimer(t *testing.T) {
	const goroutines = 1
	var wg sync.WaitGroup
	wg.Add(goroutines)
	var mu sync.Mutex
	ticker := time.NewTimer(time.Second * 3) // time.After == time.NewTimer.C

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for c := 0; c < 1000; c++ {
				mu.Lock()
				time.Sleep(time.Second * 1)
				fmt.Println("cccccc:", c)
				mu.Unlock()
			}
		}()
	}

	// 读出通道方式一【推荐】
	for {
		if v, ok := <-ticker.C; ok {
			fmt.Println(">>>>>>>>", v, ok)
			// return // 测试ticker
		}
	}
	// 读出通道方式二 // 此处只有一个通道等待，推荐使用方式一
	// for {
	//     select {
	//     case <-ticker:
	//         fmt.Println(">>>>>>>>")
	//         // return // 测试ticker
	//     }
	// }

	wg.Wait()
}

func TestToUnix1(t *testing.T) {
	var v string
	v = "2023-05-18 16:17:18"
	fmt.Println(">>>")
	fmt.Println(ToLocalUnix(v, consts.TimeLayoutV1))

	v = "2023-05-18 16:17"
	fmt.Println(">>>")
	fmt.Println(ToLocalUnix(v, consts.TimeLayoutV2))

	v = "2023-05-18 16"
	fmt.Println(">>>")
	fmt.Println(ToLocalUnix(v, consts.TimeLayoutV3))

	v = "2023-05-18"
	fmt.Println(">>>")
	fmt.Println(ToLocalUnix(v, consts.TimeLayoutV4))
}

func TestToLocalUnix(t *testing.T) {
	type args struct {
		timeStr    string
		timeLayout string
	}
	tests := []struct {
		name         string
		args         args
		wantTimeUnix int64
		wantErr      bool
	}{
		{
			name: "",
			args: args{
				timeStr:    "2023-05-18 16:17:18",
				timeLayout: consts.TimeLayoutV1,
			},
			wantTimeUnix: 0,
			wantErr:      false,
		},
		{
			name: "",
			args: args{
				timeStr:    "2023-05-18 16:17",
				timeLayout: consts.TimeLayoutV2,
			},
			wantTimeUnix: 0,
			wantErr:      false,
		},
		{
			name: "",
			args: args{
				timeStr:    "2023-05-18 16",
				timeLayout: consts.TimeLayoutV3,
			},
			wantTimeUnix: 0,
			wantErr:      false,
		},
		{
			name: "",
			args: args{
				timeStr:    "2023-05-18",
				timeLayout: consts.TimeLayoutV4,
			},
			wantTimeUnix: 0,
			wantErr:      false,
		},
		{
			name: "",
			args: args{
				timeStr:    "05-18-2023",
				timeLayout: consts.TimeLayoutV5,
			},
			wantTimeUnix: 0,
			wantErr:      false,
		},
		{
			name: "",
			args: args{
				timeStr:    "05-18-23",
				timeLayout: consts.TimeLayoutV6,
			},
			wantTimeUnix: 0,
			wantErr:      false,
		},
		{
			name: "",
			args: args{
				timeStr:    "2023-05-18",
				timeLayout: consts.TimeLayoutV4,
			},
			wantTimeUnix: 0,
			wantErr:      false,
		},
		{
			name: "",
			args: args{
				timeStr:    "05-18-23",
				timeLayout: consts.TimeLayoutV6,
			},
			wantTimeUnix: 0,
			wantErr:      false,
		},
		{
			name: "",
			args: args{
				timeStr:    "2023/05/18",
				timeLayout: consts.TimeLayoutV41,
			},
			wantTimeUnix: 0,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTimeUnix, err := ToLocalUnix(tt.args.timeStr, tt.args.timeLayout)
			fmt.Println(gotTimeUnix, err)
		})
	}
}
