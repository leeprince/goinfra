package wordcount

/*
 * @Date: 2020-09-03 17:10:58
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2020-09-03 18:07:23
 */

import (
	"github.com/leeprince/goinfra/utils/flowutil/gd/gostreaming"
	"strings"
)

var _ gostreaming.DataStreamInterface = (*WordSpliter)(nil)

type WordSpliter struct {
	*gostreaming.DataStream
}

func NewWordSpliter() gostreaming.DataStreamInterface {
	return &WordSpliter{
		DataStream: gostreaming.NewDataStream(),
	}
}

func (w *WordSpliter) Process(_ gostreaming.StatusStorage, ch <-chan *gostreaming.Event) {
	for {
		select {
		case event := <-ch:
			line := event.Data.(string)
			words := strings.Split(line, " ")

			for _, word := range words {
				word = strings.TrimSpace(word)
				if word == "" {
					continue
				}
				w.Send(&gostreaming.Event{
					Data: word,
				})
			}
		}
	}
}
