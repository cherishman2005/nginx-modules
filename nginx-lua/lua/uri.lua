local uri = "http://example.com/path/to/page?param1=value1"

-- 添加的参数
local additional_param = "param2=value2"

-- 分割 URI，将路径和查询字符串部分分开
local path, query = uri:match("([^?]+)%?(.*)")

if path and query then
    -- 将额外参数追加到原查询字符串后面
    local new_query = query .. "&" .. additional_param
    
    -- 最终的 URI 包含原始路径和新的查询字符串
    local new_uri = path .. "?" .. new_query
    print(new_uri)  -- 打印新的 URI
end