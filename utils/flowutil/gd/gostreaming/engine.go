package gostreaming

/*
 * @Date: 2020-07-06 13:37:44
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2020-07-17 09:58:55
 */

import (
	"context"
	"errors"
)

type Engine struct {
	ctx    context.Context
	cancel context.CancelFunc

	// core components
	dataSource    DataSourceInterface
	dataStreams   []DataStreamInterface
	statusStorage StatusStorage
}

func New(s StatusStorage) *Engine {
	ctx, cancel := context.WithCancel(context.Background())
	return &Engine{
		ctx:    ctx,
		cancel: cancel,

		dataSource:    nil,
		dataStreams:   make([]DataStreamInterface, 0),
		statusStorage: s,
	}
}

func (e *Engine) SetDataSource(dataSource DataSourceInterface) *Engine {
	e.dataSource = dataSource
	return e
}

func (e *Engine) AddDataStream(dataStream DataStreamInterface) *Engine {
	e.dataStreams = append(e.dataStreams, dataStream)
	return e
}

func (e *Engine) flightCheck() error {
	if e.dataSource == nil {
		return errors.New("data source can not be nil")
	}

	return nil
}

func (e *Engine) Run() error {
	err := e.flightCheck()
	if err != nil {
		return err
	}

	ch := e.dataSource.Delivery()
	for i := 0; i < len(e.dataStreams); i++ {
		for j := 0; j < e.dataStreams[i].Parallel(); j++ {
			go e.dataStreams[i].Process(e.statusStorage, ch)
		}
		ch = e.dataStreams[i].Delivery()
	}
	e.dropFinalOuputChan(ch)
	go e.dataSource.Start()

	// bloking until calling Engine.Stop()
	select {
	case <-e.ctx.Done():
		return nil
	}
}

func (e *Engine) dropFinalOuputChan(ch <-chan *Event) {
	go func(finalOutputChan <-chan *Event) {
		for {
			select {
			case <-finalOutputChan:
				// do nothing
			case <-e.ctx.Done():
				// exit goroutine
				return
			}
		}
	}(ch)
}

func (e *Engine) Stop() {
	e.dataSource.Stop()
	e.cancel()
}
