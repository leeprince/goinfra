#!/bin/sh

URL="http://localhost:19999/check-rick-control"

# 使用方法: wrk <选项> <被测 HTTP 服务的 URL>
# wrk <-c 跟服务器建立并保持的 TCP 连接数量>  <-d 压测时间，支持时间单位 (2s, 2m, 2h)>  <-t 使用多少个线程进行压测> <--latency> <--script 指定 lua 脚本路径> ${URL} <被测 HTTP 服务的 URL>
wrk -c 200 -d 2s -t 8 --latency --script ./luascript/check_risk_control.lua ${URL}