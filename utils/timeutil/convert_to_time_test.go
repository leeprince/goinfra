package timeutil

import (
	"fmt"
	"github.com/leeprince/goinfra/consts"
	"testing"
	"time"
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

func TestToCurrDayTimeByUnix(t *testing.T) {
	type args struct {
		unixTime int64
	}
	tests := []struct {
		name          string
		args          args
		wantStartTime time.Time
		wantEndTime   time.Time
	}{
		{
			name: "1",
			args: args{
				unixTime: 1698202800, // 2023-10-25 11:00:00
			},
		},
		{
			name: "2",
			args: args{
				unixTime: 1698166800, // 2023-10-25 01:00:00
			},
		},
		{
			name: "3",
			args: args{
				unixTime: 1698246000, // 2023-10-25 23:00:00
			},
		},
		{
			name: "11",
			args: args{
				unixTime: 1698116400, // 2023-10-24 11:00:00
			},
		},
		{
			name: "12",
			args: args{
				unixTime: 1698080400, // 2023-10-24 01:00:00
			},
		},
		{
			name: "13",
			args: args{
				unixTime: 1698159600, // 2023-10-24 23:00:00
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStartTime, gotEndTime := ToCurrDayTimeByUnix(tt.args.unixTime)
			fmt.Println(gotStartTime.Format(consts.TimeYYmdHis), gotEndTime.Format(consts.TimeYYmdHis))
		})
	}
}

func TestToDateTimeByStr(t *testing.T) {
	type args struct {
		timeStr    string
		timeLayout string
	}
	tests := []struct {
		name          string
		args          args
		wantStartTime time.Time
		wantEndTime   time.Time
		wantErr       bool
	}{
		{
			name: "",
			args: args{
				timeStr:    "2024-01-10 18:07:01",
				timeLayout: consts.TimeYYmdHis,
			},
			wantStartTime: time.Time{},
			wantEndTime:   time.Time{},
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStartTime, gotEndTime, err := ToDateTimeByStr(tt.args.timeStr, tt.args.timeLayout)
			if err != nil {
				fmt.Println("err:", err)
				return
			}
			fmt.Println("gotStartTime:", gotStartTime.Format(consts.TimeYYmdHis))
			fmt.Println("gotEndTime:", gotEndTime.Format(consts.TimeYYmdHis))

		})
	}
}
