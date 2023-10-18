package gostreaming

/*
 * @Date: 2020-07-06 11:59:40
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2020-07-10 14:53:21
 */

type DataSourceInterface interface {
	Start()
	Stop()

	SetName(string)
	Print()
	Send(*Event)
	Delivery() <-chan *Event
}
