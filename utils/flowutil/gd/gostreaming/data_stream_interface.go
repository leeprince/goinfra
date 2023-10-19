package gostreaming

/*
 * @Date: 2020-07-06 12:35:58
 * @LastEditors: aiden.deng(Zhenpeng Deng) 2020-07-06 17:29:07
 */

type DataStreamInterface interface {
	Process(StatusStorage, <-chan *Event)

	SetName(string)
	Name() string

	SetParallel(int)
	Parallel() int

	Send(*Event)
	Print()
	Delivery() <-chan *Event
}
