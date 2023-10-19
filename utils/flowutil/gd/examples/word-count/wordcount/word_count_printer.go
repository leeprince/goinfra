package wordcount

/*
 * @Date: 2020-09-03 17:44:10
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2020-09-04 16:55:12
 */
import (
	"fmt"
	"github.com/leeprince/goinfra/utils/flowutil/gd/gostreaming"
	"time"

	"github.com/emirpasic/gods/sets/treeset"
)

var _ gostreaming.DataStreamInterface = (*WordCountPrinter)(nil)

type WordCountPrinter struct {
	*gostreaming.DataStream
}

func NewWordCountPrinter() gostreaming.DataStreamInterface {
	return &WordCountPrinter{
		DataStream: gostreaming.NewDataStream(),
	}
}

func (w *WordCountPrinter) Process(statusStorage gostreaming.StatusStorage, ch <-chan *gostreaming.Event) {
	ticker := time.NewTicker(1 * time.Second)
	wordSet := treeset.NewWithStringComparator()
	iterTime := 0

	for {
		select {
		case event := <-ch:
			word := event.Data.(string)
			wordSet.Add(word)
			w.Send(event)
		case <-ticker.C:
			words, counts := w.getWordCounts(statusStorage, wordSet)
			if len(words) > 0 {
				fmt.Println("")
			}
			for i := range words {
				fmt.Printf("[INFO] [word-count-printer] [%d] %s: %d \n", iterTime, words[i], counts[i])
			}

			iterTime++

		}
	}
}

func (w *WordCountPrinter) getWordCounts(statusStorage gostreaming.StatusStorage, wordSet *treeset.Set) ([]string, []int) {
	batch := gostreaming.NewBatch()

	words := make([]string, wordSet.Size())
	counts := make([]int, wordSet.Size())

	for i, wordInterface := range wordSet.Values() {
		word := wordInterface.(string)
		words[i] = word
		batch.Get([]string{word}, "word-count", nil)
	}

	results, errs, err := statusStorage.ExecBatch(batch)
	if err != nil {
		fmt.Printf("[ERROR] err: %+v \n", err)
		return nil, nil
	}

	for _, err := range errs {
		if err != nil {
			fmt.Printf("[WARN] err: %+v \n", err)
		}
	}

	for i, result := range results {
		counts[i] = result.(int)
	}
	return words, counts
}
