#!/bin/sh

URL="https://apigw-local-test.goldentec.com/file-center/format-ctl/pdf-to-file"

# 使用方法: wrk <选项> <被测 HTTP 服务的 URL>
# wrk <-c 跟服务器建立并保持的 TCP 连接数量> <-d 压测时间，支持时间单位 (2s, 2m, 2h)> <-t 使用多少个线程进行压测> <--latency> <--script 指定 lua 脚本路径> ${URL} <被测 HTTP 服务的 URL>
wrk -c 30 -d 10s -t 30 --latency --script ./luascript/pdf_to_jpeg.lua ${URL}