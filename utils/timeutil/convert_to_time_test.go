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
	fmt.Println(ToTimeUnix(v, consts.TimeYYmdHis))

	v = "2023-05-18 16:17"
	fmt.Println(">>>")
	fmt.Println(ToTimeUnix(v, consts.TimeYYmdHi))

	v = "2023-05-18 16"
	fmt.Println(">>>")
	fmt.Println(ToTimeUnix(v, consts.TimeYYmdH))

	v = "2023-05-18"
	fmt.Println(">>>")
	fmt.Println(ToTimeUnix(v, consts.TimeYYmd))

	v = "202308"
	fmt.Println(">>>")
	fmt.Println(ToTimeUnix(v, consts.TimeYYmm))
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
				timeLayout: consts.TimeYYmdHis,
			},
			wantTimeUnix: 0,
			wantErr:      false,
		},
		{
			name: "",
			args: args{
				timeStr:    "2023-05-18 16:17",
				timeLayout: consts.TimeYYmdHi,
			},
			wantTimeUnix: 0,
			wantErr:      false,
		},
		{
			name: "",
			args: args{
				timeStr:    "2023-05-18 16",
				timeLayout: consts.TimeYYmdH,
			},
			wantTimeUnix: 0,
			wantErr:      false,
		},
		{
			name: "",
			args: args{
				timeStr:    "2023-05-18",
				timeLayout: consts.TimeYYmd,
			},
			wantTimeUnix: 0,
			wantErr:      false,
		},
		{
			name: "",
			args: args{
				timeStr:    "05-18-2023",
				timeLayout: consts.TimeMdyy,
			},
			wantTimeUnix: 0,
			wantErr:      false,
		},
		{
			name: "",
			args: args{
				timeStr:    "05-18-23",
				timeLayout: consts.TimeMdy,
			},
			wantTimeUnix: 0,
			wantErr:      false,
		},
		{
			name: "",
			args: args{
				timeStr:    "2023-05-18",
				timeLayout: consts.TimeYYmd,
			},
			wantTimeUnix: 0,
			wantErr:      false,
		},
		{
			name: "",
			args: args{
				timeStr:    "05-18-23",
				timeLayout: consts.TimeMdy,
			},
			wantTimeUnix: 0,
			wantErr:      false,
		},
		{
			name: "",
			args: args{
				timeStr:    "2023/05/18",
				timeLayout: consts.TimeYymd,
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
	d1 := MonthFirthDayTime()
	fmt.Println(d1)
}

func TestGetYearFirthMonthTime(t *testing.T) {
	m1 := YearFirthMonthTime()
	fmt.Println(m1)
}

func TestGetPreMonthUnixTimeRange(t *testing.T) {
	gotStartTime, gotEndTime := PreMonthUnixTimeRange()

	fmt.Println(gotStartTime.Unix())
	fmt.Println(gotEndTime.Unix())

	fmt.Println(gotStartTime.Format("2006-01-02 15:04:05"))
	fmt.Println(gotEndTime.Format("2006-01-02 15:04:05"))
}
