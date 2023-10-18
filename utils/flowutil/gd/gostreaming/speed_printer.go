package gostreaming

/*
 * @Date: 2020-07-06 17:11:23
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2020-07-13 08:58:00
 */

import (
	"fmt"
	"sync"
	"time"
)

var _ DataStreamInterface = (*SpeedPrinter)(nil)

type SpeedPrinter struct {
	*DataStream
}

func NewSpeedPrinter(name string) *SpeedPrinter {
	d := &SpeedPrinter{
		DataStream: NewDataStream(),
	}
	d.SetName(name)
	return d
}

func (d *SpeedPrinter) Process(_ StatusStorage, ch <-chan *Event) {
	mu := sync.Mutex{}
	n := new(int)
	*n = 0

	go func(n *int) {
		for {
			time.Sleep(1 * time.Second)
			mu.Lock()
			fmt.Printf("[gostreaming] [%s] [%s] [%s]: %s\n",
				d.Name(),
				time.Now().Format("2006-01-02 15:04:05"),
				"INFO",
				fmt.Sprintf("n/sec: %d", *n),
			)
			*n = 0
			mu.Unlock()
		}
	}(n)

	for {
		select {
		case event := <-ch:
			mu.Lock()
			*n = *n + 1
			mu.Unlock()
			d.Send(event)
		}
	}
}
