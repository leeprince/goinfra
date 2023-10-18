package gostreaming

import "fmt"

/*
 * @Date: 2020-07-06 13:56:53
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2020-07-10 15:10:17
 */

var _ DataSourceInterface = (*DataSource)(nil)

type DataSource struct {
	name     string
	ch       chan *Event
	usePrint bool
}

func NewDataSource() *DataSource {
	dataSource := &DataSource{
		ch: make(chan *Event),
	}
	return dataSource
}

func (d *DataSource) Start() {
	panic(fmt.Sprintf("DataSource(name:\"%s\") did not implement DataSource.Start()", d.name))
}

func (d *DataSource) Stop() {
	panic(fmt.Sprintf("DataSource(name:\"%s\") did not implement DataSource.Start()", d.name))
}

func (d *DataSource) SetName(name string) {
	d.name = name
}

func (d *DataSource) Print() {
	d.usePrint = true
}

func (d *DataSource) Send(event *Event) {
	if d.usePrint {
		fmt.Printf("[gostreaming] [INFO] [DataSource:%s]: event.Data: %+v\n", d.name, event.Data)
	}
	d.ch <- event
}

func (d *DataSource) Delivery() <-chan *Event {
	return d.ch
}
