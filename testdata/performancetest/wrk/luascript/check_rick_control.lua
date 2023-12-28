local wrk = wrk
local require = require
local tostring = tostring
local math = require('math')
local os = require('os')
local cjson = require('cjson')
local os_time = os.time

math.randomseed(os_time())

function request()
    -- 使用单引号，包含json对象则里面的双引号不会被转义的
    -- local body = '{"open_org_id":"p-open_org_id","transaction_type":"01","transaction_sn":"p-transaction_sn","transaction_amount":200,"payer_account_no":"p-payer_account_no","payer_account_name":"p-payer_account_name","payee_account_no":"p-payee_account_no","payee_account_name":"p-payee_account_name"}'
    -- 使用随机数
    local open_org_id = math.random(1, 15955)
    local body = '{"open_org_id":"'..tostring(open_org_id)..'","transaction_type":"01","transaction_sn":"p-transaction_sn","transaction_amount":200,"payer_account_no":"p-payer_account_no","payer_account_name":"p-payer_account_name","payee_account_no":"p-payee_account_no","payee_account_name":"p-payee_account_name"}'

    return wrk.format('POST', nil, nil, body)
end

function response(status, headers, body)
    rsp = cjson.decode(body)
    if rsp.code ~= 0 then
        return
    end
    -- os.exit(1)
end