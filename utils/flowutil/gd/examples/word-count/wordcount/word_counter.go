package wordcount

/*
 * @Date: 2020-09-03 17:10:58
 * @LastEditors: aiden.deng(Zhenpeng Deng)
 * @LastEditTime: 2021-04-15 16:20:50
 */

import (
	"time"

	"gitlab.yewifi.com/risk-control/risk-common/pkg/gostreaming"
)

var _ gostreaming.DataStreamInterface = (*WordCounter)(nil)

type WordCounter struct {
	*gostreaming.DataStream
}

func NewWordCounter() gostreaming.DataStreamInterface {
	return &WordCounter{
		DataStream: gostreaming.NewDataStream(),
	}
}

func (w *WordCounter) Process(statusStorage gostreaming.StatusStorage, ch <-chan *gostreaming.Event) {
	for {
		select {
		// 监听事件
		case event := <-ch:
			word := event.Data.(string)
			batch := gostreaming.NewBatch()

			// args: primaryKeys, targetName, descriptions
			// => gs_<primaryKey[0]>_..._<primaryKey[n-1]>_(targetName)_[descriptions[0]]_...[descriptions[n-1]]
			batch.Incr([]string{word}, "word-count", nil, 30*time.Second)
			statusStorage.ExecBatch(batch)

			w.Send(event)
		}
	}
}
