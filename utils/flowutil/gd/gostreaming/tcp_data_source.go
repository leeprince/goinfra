package gostreaming

/*
 * @Date: 2020-09-03 16:57:28
 * @LastEditors: aiden.deng(Zhenpeng Deng)
 * @LastEditTime: 2020-09-03 17:10:01
 */

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

var _ DataSourceInterface = (*TCPDataSource)(nil)

type TCPDataSource struct {
	*DataSource
	lis       net.Listener
	isRunning bool
}

func MustNewTCPDataSource(port int) DataSourceInterface {
	ds, err := NewTCPDataSource(port)
	if err != nil {
		panic(err)
	}
	return ds
}

func NewTCPDataSource(port int) (DataSourceInterface, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	ds := &TCPDataSource{
		DataSource: NewDataSource(),
		lis:        lis,
		isRunning:  false,
	}
	return ds, nil
}

func (ds *TCPDataSource) Start() {
	for !ds.isRunning {
		conn, err := ds.lis.Accept()
		if err != nil {
			fmt.Printf("[ERROR] err %+v", err)
			continue
		}
		go ds.handleConn(conn)
	}
}

func (ds *TCPDataSource) handleConn(conn net.Conn) {
	rd := bufio.NewReader(conn)

	for {
		data, _, err := rd.ReadLine()
		if err != nil {
			if err == io.EOF {
				fmt.Println("[INFO] EOF")
				return
			}
			fmt.Printf("[ERROR] err: %+v", err)
			return
		}
		ds.Send(&Event{
			Data: string(data),
		})
	}
}

func (ds *TCPDataSource) Stop() {
	ds.isRunning = false
}
