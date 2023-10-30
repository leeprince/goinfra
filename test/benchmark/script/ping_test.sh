#!/bin/bash

URL="http://localhost:19999/ping"

wrk -c 200 -t 8 -d 2 --latency --script ./lua/ping.lua ${URL}