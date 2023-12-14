
local wrk = wrk
local require = require
local table_concat = table.concat

wrk.method = "POST"
wrk.headers["Content-Type"] = "application/type"
wrk.headers["Authorization"] = ""

wrk.body = '{"echo":"hello world"}'