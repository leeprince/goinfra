package gostreaming

/*
 * @Date: 2020-07-06 13:48:48
 * @LastEditors: aiden.deng(Zhenpeng Deng)
 * @LastEditTime: 2020-07-10 15:07:31
 */

import (
	"bufio"
	"io"
	"os"
)

var _ DataSourceInterface = (*FileDataSource)(nil)

type FileDataSource struct {
	*DataSource

	isRunning bool
	rd        *bufio.Reader
}

func MustNewFileDataSource(filePath string) DataSourceInterface {
	fds, err := NewFileDataSource(filePath)
	if err != nil {
		panic(err)
	}
	return fds
}

func NewFileDataSource(filePath string) (DataSourceInterface, error) {
	fp, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	fds := &FileDataSource{
		DataSource: NewDataSource(),

		isRunning: false,
		rd:        bufio.NewReader(fp),
	}
	return fds, nil
}

func (fds *FileDataSource) Start() {
	for !fds.isRunning {
		data, _, err := fds.rd.ReadLine()
		if err != nil {
			if err == io.EOF {
				return
			}
			panic(err)
		}
		event := &Event{
			Data: string(data),
		}
		fds.Send(event)
	}
}

func (fds *FileDataSource) Stop() {
	fds.isRunning = false
}
