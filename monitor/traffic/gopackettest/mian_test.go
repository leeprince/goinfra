package main

import "testing"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/5/7 23:26
 * @Desc:
 */

func TestMacIfconfig(t *testing.T) {
	MacIfconfig()
}

func TestMonitorLoopback(t *testing.T) {
	MonitorMACLoopback()
}

func TestMonitorLoopbackNoserver(t *testing.T) {
	MonitorMACLoopbackNoserver()
}

func TestMonitorMACLoopbackNoserverNoclient(t *testing.T) {
	MonitorMACLoopbackNoserverNoclient()
}
