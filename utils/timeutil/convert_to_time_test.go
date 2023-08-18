package timeutil

import (
	"fmt"
	"github.com/leeprince/goinfra/consts"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/15 17:45
 * @Desc:
 */

func TestToTimeUnix(t *testing.T) {
	var v string
	v = "2023-05-18 16:17:18"
	fmt.Println(">>>")
	fmt.Println(ToTimeUnix(v, consts.TimeLayoutV1))

	v = "2023-05-18 16:17"
	fmt.Println(">>>")
	fmt.Println(ToTimeUnix(v, consts.TimeLayoutV2))

	v = "2023-05-18 16"
	fmt.Println(">>>")
	fmt.Println(ToTimeUnix(v, consts.TimeLayoutV3))

	v = "2023-05-18"
	fmt.Println(">>>")
	fmt.Println(ToTimeUnix(v, consts.TimeLayoutV4))

	v = "202308"
	fmt.Println(">>>")
	fmt.Println(ToTimeUnix(v, consts.TimeLayoutV71))
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
			gotTimeUnix, err := ToTimeUnix(tt.args.timeStr, tt.args.timeLayout)
			fmt.Println(gotTimeUnix, err)
		})
	}
}

func TestUnixToTime(t *testing.T) {
	timeUnix := int64(1622517600)
	timeObj := ToTimeByUnix(timeUnix)
	timeUnixMonthString := MonthT(timeObj)
	fmt.Println(timeUnixMonthString)

}

func TestGetMonthFirthDay(t *testing.T) {
	d1 := GetMonthFirthDayTime()
	fmt.Println(d1)
}

func TestGetYearFirthMonthTime(t *testing.T) {
	m1 := GetYearFirthMonthTime()
	fmt.Println(m1)
}
