local wrk = wrk
local require = require
local tostring = tostring
local math = require('math')
local os = require('os')
local cjson = require('cjson')
local os_time = os.time

math.randomseed(os_time())

function request()
    -- 定义一个列表
    local list = {"https://kpserverdev-1251506165.cos.ap-shanghai.myqcloud.com/wbx/upload/OQgorPo0MckWeb374dae2b7e56f9ccb9a6ea1ad0d276_20201844_1703557617396552.pdf",
                  "https://kpserverdev-1251506165.cos.ap-shanghai.myqcloud.com/e-document-import-ctl/test/0001.pdf",
                  "https://kpserverdev-1251506165.cos.ap-shanghai.myqcloud.com/e-invoice-invoice-ctl/test/00193a3ffe9275b2_7095231768642499417.pdf",
                  "https://kpserverdev-1251506165.cos.ap-shanghai.myqcloud.com/e-invoice-invoice-ctl/test/01b6aca9fa10190c_7107914276244862765.pdf",
                  "https://kpserverdev-1251506165.cos.ap-shanghai.myqcloud.com/e-document-import-ctl/local/231aaadabb31b2ce-0-22442000000006722172.pdf"}

    -- 随机选择一个元素
    local randomElement = list[math.random(#list)]

    -- 打印
    --print(randomElement)

    local body = '{"pdf_url":"'..randomElement..'"}'

    return wrk.format('POST', nil, nil, body)
end

--[[
function response(status, headers, body)
    rsp = cjson.decode(body)
    if rsp.code ~= 0 then
        return
    end
    -- os.exit(1)
end]]
