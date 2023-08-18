package timeutil

import (
	"fmt"
	"github.com/spf13/cast"
	"testing"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/15 15:40
 * @Desc:
 */

func TestGetBeforeMonth(t *testing.T) {
	fmt.Println(GetHistoryMonth(1), "---(1)--", GetHistoryMonth(1).Format("200601"))
	fmt.Println(GetHistoryMonth(2), "---(2)--", GetHistoryMonth(2).Format("200601"))
	fmt.Println(GetHistoryMonth(3), "---(3)--", GetHistoryMonth(3).Format("200601"))
	fmt.Println(GetHistoryMonth(4), "---(4)--", GetHistoryMonth(4).Format("200601"))
	fmt.Println(GetHistoryMonth(5), "---(5)--", GetHistoryMonth(5).Format("200601"))
	fmt.Println(GetHistoryMonth(6), "---(6)--", GetHistoryMonth(6).Format("200601"))
	fmt.Println(GetHistoryMonth(7), "---(7)--", GetHistoryMonth(7).Format("200601"))
	fmt.Println(GetHistoryMonth(8), "---(8)--", GetHistoryMonth(8).Format("200601"))
	fmt.Println(GetHistoryMonth(9), "---(9)--", GetHistoryMonth(9).Format("200601"))
	fmt.Println(GetHistoryMonth(10), "---(10)--", GetHistoryMonth(10).Format("200601"))
	fmt.Println(GetHistoryMonth(11), "---(11)--", GetHistoryMonth(11).Format("200601"))
	fmt.Println(GetHistoryMonth(12), "---(12)--", GetHistoryMonth(12).Format("200601"))
	fmt.Println(GetHistoryMonth(13), "---(13)--", GetHistoryMonth(13).Format("200601"))
}

func TestGetTimeBeforeMonth(t *testing.T) {
	myMonth := 202308

	useTime := cast.ToTime(myMonth)
	useTimeString := useTime.Format("2006-01-02 15:04:05")
	fmt.Println("useTimeString 有问题的:", useTimeString)

	l := time.Local
	useTime = time.Date(2023, 8, 1, 0, 0, 0, 0, l)
	useTimeString = useTime.Format("2006-01-02 15:04:05")
	fmt.Println("useTimeString 正确的:", useTimeString)

	fmt.Println(GetHistoryMonthByTime(useTime, 1), "---(1)--", GetHistoryMonthByTime(useTime, 1).Format("200601"))
	fmt.Println(GetHistoryMonthByTime(useTime, 2), "---(2)--", GetHistoryMonthByTime(useTime, 2).Format("200601"))
	fmt.Println(GetHistoryMonthByTime(useTime, 3), "---(3)--", GetHistoryMonthByTime(useTime, 3).Format("200601"))
	fmt.Println(GetHistoryMonthByTime(useTime, 4), "---(4)--", GetHistoryMonthByTime(useTime, 4).Format("200601"))
	fmt.Println(GetHistoryMonthByTime(useTime, 5), "---(5)--", GetHistoryMonthByTime(useTime, 5).Format("200601"))
	fmt.Println(GetHistoryMonthByTime(useTime, 6), "---(6)--", GetHistoryMonthByTime(useTime, 6).Format("200601"))
	fmt.Println(GetHistoryMonthByTime(useTime, 7), "---(7)--", GetHistoryMonthByTime(useTime, 7).Format("200601"))
	fmt.Println(GetHistoryMonthByTime(useTime, 8), "---(8)--", GetHistoryMonthByTime(useTime, 8).Format("200601"))
	fmt.Println(GetHistoryMonthByTime(useTime, 9), "---(9)--", GetHistoryMonthByTime(useTime, 9).Format("200601"))
	fmt.Println(GetHistoryMonthByTime(useTime, 10), "---(10)--", GetHistoryMonthByTime(useTime, 10).Format("200601"))
	fmt.Println(GetHistoryMonthByTime(useTime, 11), "---(11)--", GetHistoryMonthByTime(useTime, 11).Format("200601"))
	fmt.Println(GetHistoryMonthByTime(useTime, 12), "---(12)--", GetHistoryMonthByTime(useTime, 12).Format("200601"))
	fmt.Println(GetHistoryMonthByTime(useTime, 13), "---(13)--", GetHistoryMonthByTime(useTime, 13).Format("200601"))
}

func TestGetTimeBeforeMonthOfUseMonth(t *testing.T) {
	useYear := 2022
	useMonth := 8

	fmt.Println(GetHistoryMonthByUseMonth(useYear, useMonth, 1), "---(1)--", GetHistoryMonthByUseMonth(useYear, useMonth, 1).Format("200601"))
	fmt.Println(GetHistoryMonthByUseMonth(useYear, useMonth, 2), "---(2)--", GetHistoryMonthByUseMonth(useYear, useMonth, 2).Format("200601"))
	fmt.Println(GetHistoryMonthByUseMonth(useYear, useMonth, 3), "---(3)--", GetHistoryMonthByUseMonth(useYear, useMonth, 3).Format("200601"))
	fmt.Println(GetHistoryMonthByUseMonth(useYear, useMonth, 4), "---(4)--", GetHistoryMonthByUseMonth(useYear, useMonth, 4).Format("200601"))
	fmt.Println(GetHistoryMonthByUseMonth(useYear, useMonth, 5), "---(5)--", GetHistoryMonthByUseMonth(useYear, useMonth, 5).Format("200601"))
	fmt.Println(GetHistoryMonthByUseMonth(useYear, useMonth, 6), "---(6)--", GetHistoryMonthByUseMonth(useYear, useMonth, 6).Format("200601"))
	fmt.Println(GetHistoryMonthByUseMonth(useYear, useMonth, 7), "---(7)--", GetHistoryMonthByUseMonth(useYear, useMonth, 7).Format("200601"))
	fmt.Println(GetHistoryMonthByUseMonth(useYear, useMonth, 8), "---(8)--", GetHistoryMonthByUseMonth(useYear, useMonth, 8).Format("200601"))
	fmt.Println(GetHistoryMonthByUseMonth(useYear, useMonth, 9), "---(9)--", GetHistoryMonthByUseMonth(useYear, useMonth, 9).Format("200601"))
	fmt.Println(GetHistoryMonthByUseMonth(useYear, useMonth, 10), "---(10)--", GetHistoryMonthByUseMonth(useYear, useMonth, 10).Format("200601"))
	fmt.Println(GetHistoryMonthByUseMonth(useYear, useMonth, 11), "---(11)--", GetHistoryMonthByUseMonth(useYear, useMonth, 11).Format("200601"))
	fmt.Println(GetHistoryMonthByUseMonth(useYear, useMonth, 12), "---(12)--", GetHistoryMonthByUseMonth(useYear, useMonth, 12).Format("200601"))
	fmt.Println(GetHistoryMonthByUseMonth(useYear, useMonth, 13), "---(13)--", GetHistoryMonthByUseMonth(useYear, useMonth, 13).Format("200601"))
}

func TestToHistoryDataList(t *testing.T) {
	type args struct {
		date     string
		oldMonth int
	}
	tests := []struct {
		name          string
		args          args
		wantMonthList []int64
		wantErr       bool
	}{
		{
			name: "",
			args: args{
				date:     "202308",
				oldMonth: 12,
			},
			wantMonthList: nil,
			wantErr:       false,
		},
		{
			name: "",
			args: args{
				date:     "202308",
				oldMonth: 2,
			},
			wantMonthList: nil,
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMonthList, err := GetHistoryMonthListByMonth(tt.args.date, tt.args.oldMonth)
			fmt.Println(err)
			fmt.Println(gotMonthList)
		})
	}
}
