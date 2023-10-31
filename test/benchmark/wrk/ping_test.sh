#!/bin/bash

URL="http://localhost:19999/ping"

wrk -c 200 -d 2 -t 8--latency --script ./lua/ping.lua ${URL}