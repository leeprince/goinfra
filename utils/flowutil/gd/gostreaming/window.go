package gostreaming

/*
 * @Date: 2020-07-08 14:16:07
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2020-07-11 14:33:23
 */

import (
	"fmt"
	"time"
)

type Window struct {
	start time.Time
	end   time.Time
}

func NewWindow(start time.Time, end time.Time) *Window {
	return &Window{
		start: start,
		end:   end,
	}
}

func NewWindowByStartStartAndIntervalDay(startTime time.Time, day int) *Window {
	return &Window{
		start: startTime,
		end:   startTime.AddDate(0, 0, day-1),
	}
}

func NewWindowByEndTimeAndIntervalDay(endTime time.Time, day int) *Window {
	return &Window{
		start: endTime.AddDate(0, 0, -day+1),
		end:   endTime,
	}
}

func (w *Window) String() string {
	return fmt.Sprintf("<Window start=%s end=%s", w.start, w.end)
}

func (w *Window) ToDateStrings() []string {
	dateStrings := make([]string, 0)
	endDateString := w.end.Format("2006-01-02")
	i := 0
	for {
		curEnd := w.start.AddDate(0, 0, i)
		curEndDateString := curEnd.Format("2006-01-02")
		if curEndDateString > endDateString {
			break
		}
		dateStrings = append(dateStrings, curEndDateString)
		i++
	}
	return dateStrings
}
