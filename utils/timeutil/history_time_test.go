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
	fmt.Println(HistoryMonth(1), "---(1)--", HistoryMonth(1).Format("200601"))
	fmt.Println(HistoryMonth(2), "---(2)--", HistoryMonth(2).Format("200601"))
	fmt.Println(HistoryMonth(3), "---(3)--", HistoryMonth(3).Format("200601"))
	fmt.Println(HistoryMonth(4), "---(4)--", HistoryMonth(4).Format("200601"))
	fmt.Println(HistoryMonth(5), "---(5)--", HistoryMonth(5).Format("200601"))
	fmt.Println(HistoryMonth(6), "---(6)--", HistoryMonth(6).Format("200601"))
	fmt.Println(HistoryMonth(7), "---(7)--", HistoryMonth(7).Format("200601"))
	fmt.Println(HistoryMonth(8), "---(8)--", HistoryMonth(8).Format("200601"))
	fmt.Println(HistoryMonth(9), "---(9)--", HistoryMonth(9).Format("200601"))
	fmt.Println(HistoryMonth(10), "---(10)--", HistoryMonth(10).Format("200601"))
	fmt.Println(HistoryMonth(11), "---(11)--", HistoryMonth(11).Format("200601"))
	fmt.Println(HistoryMonth(12), "---(12)--", HistoryMonth(12).Format("200601"))
	fmt.Println(HistoryMonth(13), "---(13)--", HistoryMonth(13).Format("200601"))
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

	fmt.Println(HistoryMonthByTime(useTime, 1), "---(1)--", HistoryMonthByTime(useTime, 1).Format("200601"))
	fmt.Println(HistoryMonthByTime(useTime, 2), "---(2)--", HistoryMonthByTime(useTime, 2).Format("200601"))
	fmt.Println(HistoryMonthByTime(useTime, 3), "---(3)--", HistoryMonthByTime(useTime, 3).Format("200601"))
	fmt.Println(HistoryMonthByTime(useTime, 4), "---(4)--", HistoryMonthByTime(useTime, 4).Format("200601"))
	fmt.Println(HistoryMonthByTime(useTime, 5), "---(5)--", HistoryMonthByTime(useTime, 5).Format("200601"))
	fmt.Println(HistoryMonthByTime(useTime, 6), "---(6)--", HistoryMonthByTime(useTime, 6).Format("200601"))
	fmt.Println(HistoryMonthByTime(useTime, 7), "---(7)--", HistoryMonthByTime(useTime, 7).Format("200601"))
	fmt.Println(HistoryMonthByTime(useTime, 8), "---(8)--", HistoryMonthByTime(useTime, 8).Format("200601"))
	fmt.Println(HistoryMonthByTime(useTime, 9), "---(9)--", HistoryMonthByTime(useTime, 9).Format("200601"))
	fmt.Println(HistoryMonthByTime(useTime, 10), "---(10)--", HistoryMonthByTime(useTime, 10).Format("200601"))
	fmt.Println(HistoryMonthByTime(useTime, 11), "---(11)--", HistoryMonthByTime(useTime, 11).Format("200601"))
	fmt.Println(HistoryMonthByTime(useTime, 12), "---(12)--", HistoryMonthByTime(useTime, 12).Format("200601"))
	fmt.Println(HistoryMonthByTime(useTime, 13), "---(13)--", HistoryMonthByTime(useTime, 13).Format("200601"))
}

func TestGetTimeBeforeMonthOfUseMonth(t *testing.T) {
	useYear := 2022
	useMonth := 8

	fmt.Println(HistoryMonthByUseMonth(useYear, useMonth, 1), "---(1)--", HistoryMonthByUseMonth(useYear, useMonth, 1).Format("200601"))
	fmt.Println(HistoryMonthByUseMonth(useYear, useMonth, 2), "---(2)--", HistoryMonthByUseMonth(useYear, useMonth, 2).Format("200601"))
	fmt.Println(HistoryMonthByUseMonth(useYear, useMonth, 3), "---(3)--", HistoryMonthByUseMonth(useYear, useMonth, 3).Format("200601"))
	fmt.Println(HistoryMonthByUseMonth(useYear, useMonth, 4), "---(4)--", HistoryMonthByUseMonth(useYear, useMonth, 4).Format("200601"))
	fmt.Println(HistoryMonthByUseMonth(useYear, useMonth, 5), "---(5)--", HistoryMonthByUseMonth(useYear, useMonth, 5).Format("200601"))
	fmt.Println(HistoryMonthByUseMonth(useYear, useMonth, 6), "---(6)--", HistoryMonthByUseMonth(useYear, useMonth, 6).Format("200601"))
	fmt.Println(HistoryMonthByUseMonth(useYear, useMonth, 7), "---(7)--", HistoryMonthByUseMonth(useYear, useMonth, 7).Format("200601"))
	fmt.Println(HistoryMonthByUseMonth(useYear, useMonth, 8), "---(8)--", HistoryMonthByUseMonth(useYear, useMonth, 8).Format("200601"))
	fmt.Println(HistoryMonthByUseMonth(useYear, useMonth, 9), "---(9)--", HistoryMonthByUseMonth(useYear, useMonth, 9).Format("200601"))
	fmt.Println(HistoryMonthByUseMonth(useYear, useMonth, 10), "---(10)--", HistoryMonthByUseMonth(useYear, useMonth, 10).Format("200601"))
	fmt.Println(HistoryMonthByUseMonth(useYear, useMonth, 11), "---(11)--", HistoryMonthByUseMonth(useYear, useMonth, 11).Format("200601"))
	fmt.Println(HistoryMonthByUseMonth(useYear, useMonth, 12), "---(12)--", HistoryMonthByUseMonth(useYear, useMonth, 12).Format("200601"))
	fmt.Println(HistoryMonthByUseMonth(useYear, useMonth, 13), "---(13)--", HistoryMonthByUseMonth(useYear, useMonth, 13).Format("200601"))
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
			gotMonthList, err := HistoryMonthListByMonth(tt.args.date, tt.args.oldMonth)
			fmt.Println(err)
			fmt.Println(gotMonthList)
		})
	}
}
