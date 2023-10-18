package gostreaming

/*
 * @Date: 2020-07-09 17:33:00
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2020-07-10 15:21:50
 */

import (
	"fmt"
	"time"
)

var _ DataStreamInterface = (*EventCounter)(nil)

type EventCounter struct {
	*DataStream
	duration time.Duration
}

func NewEventCounter(duration time.Duration) *EventCounter {
	d := &EventCounter{
		DataStream: NewDataStream(),
		duration:   duration,
	}
	d.SetName("event-counter")
	return d
}

func (d *EventCounter) Process(_ StatusStorage, ch <-chan *Event) {
	i := 0
	ticker := time.NewTicker(d.duration)
	startTime := time.Now()
	for {
		select {
		case event := <-ch:
			i++
			d.Send(event)
		case <-ticker.C:
			fmt.Printf("[gostreaming] [%s] [%s] [%s]: %s\n",
				d.Name(),
				time.Now().Format("2006-01-02 15:04:05"),
				"INFO",
				fmt.Sprintf("since: %ds, num: %d", int64(time.Now().Sub(startTime).Seconds()), i),
			)
		}
	}
}
