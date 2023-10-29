package gostreaming

/*
 * @Date: 2020-07-09 09:46:45
 * @LastEditors: aiden.deng(Zhenpeng Deng)
 * @LastEditTime: 2021-04-09 17:42:13
 */

import (
	"fmt"
	"time"
)

type BatchCommand struct {
	Name  string
	Key   string
	Value interface{}

	// hash
	Field string

	// zset
	Score  float64
	Member interface{}
	Min    string
	Max    string
	Offset int64
	Count  int64

	IncrBy int

	Expiration time.Duration
}

func (b BatchCommand) String() string {
	return fmt.Sprintf("<BatchCommand Name: %s Key: %s Value: %+v Field: %s Score: %f Member: %+v IncrBy: %d>",
		b.Name, b.Key, b.Value, b.Field, b.Score, b.Member, b.IncrBy)
}
