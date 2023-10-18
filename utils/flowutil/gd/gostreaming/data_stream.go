package gostreaming

/*
 * @Date: 2020-07-06 13:57:20
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2020-09-02 16:33:53
 */

import (
	"fmt"
)

var _ DataStreamInterface = (*DataStream)(nil)

type DataStream struct {
	name      string
	outChan   chan *Event
	nParallel int
	usePrint  bool
}

func NewDataStream() *DataStream {
	return &DataStream{
		name:      "",
		outChan:   make(chan *Event),
		nParallel: 1, // 并发度
		usePrint:  false,
	}
}

func (d *DataStream) Process(statuStorage StatusStorage, inChan <-chan *Event) {
	panic(fmt.Sprintf("DataStream(name:\"%s\") did not implement DataStream.Process()", d.name))
}

func (d *DataStream) SetName(name string) {
	d.name = name
}

func (d *DataStream) Name() string {
	return d.name
}

func (d *DataStream) SetParallel(n int) {
	if n < 1 {
		panic(fmt.Errorf("parallel number can not less than 1"))
	}
	d.nParallel = n
}

func (d *DataStream) Parallel() int {
	return d.nParallel
}

func (d *DataStream) Send(event *Event) {
	if d.usePrint {
		fmt.Printf("[gostreaming] [INFO] [DataStream:%s]: event.Data: %+v\n", d.name, event.Data)
	}
	d.outChan <- event
}

func (d *DataStream) Print() {
	d.usePrint = true
}

func (d *DataStream) Delivery() <-chan *Event {
	return d.outChan
}
