package gostreaming

/*
 * @Date: 2020-07-09 11:20:54
 * @LastEditors: aiden.deng(Zhenpeng Deng)
 * @LastEditTime: 2021-04-12 16:06:21
 */

import (
	"fmt"
	"strings"
)

type BatchInterface interface {
	Size() int
	GetBatchCommands() []*BatchCommand
	String() string
}

func makeKeyName(primaryKeys []string, targetName string, descriptions []string) string {
	fields := make([]string, 0)
	fields = append(fields, "gs")

	for _, primaryKey := range primaryKeys {
		fields = append(fields, fmt.Sprintf("<%s>", primaryKey))
	}

	fields = append(fields, fmt.Sprintf("(%s)", targetName))

	for _, description := range descriptions {
		fields = append(fields, fmt.Sprintf("[%s]", description))
	}

	return strings.Join(fields, ":")
}
