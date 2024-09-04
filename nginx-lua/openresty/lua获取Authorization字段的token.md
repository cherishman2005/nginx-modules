# openresty lua 请求获取Authorization: Bearer 字段的token 

在OpenResty的Lua环境中，当你想要获取请求头中Authorization: Bearer字段的token时，你需要首先使用get_header方法获取整个Authorization字段的值，然后解析这个值以提取Bearer token。通常，Bearer token会紧跟在Bearer关键字之后，并且它们之间可能有一个空格。

以下是一个Lua代码示例，展示了如何实现这一点：

```lua
location /test {
    default_type 'text/plain';
    content_by_lua_block {
        -- 获取请求对象
        local req = ngx.req

        -- 获取Authorization请求头字段
        local authorization_header = req.get_header("Authorization")

        -- 检查Authorization请求头是否存在
        if authorization_header and authorization_header:starts_with("Bearer ") then
            -- 提取Bearer token
            local token = authorization_header:sub(8) -- 从"Bearer "之后开始提取

            -- 输出token
            ngx.say("Bearer Token: ", token)
        else
            -- 如果Authorization请求头不存在或格式不正确
            ngx.say("Authorization header is not present or does not contain a Bearer token.")
        end
    }
}
```

注意：上面的代码示例使用了:starts_with和:sub这两个字符串操作方法，但它们是假设存在的，因为在标准的Lua或OpenResty的ngx.reqAPI中并不直接提供这些方法。实际上，你需要自己实现这些字符串处理功能，或者使用OpenResty提供的其他字符串处理模块（如果有的话）。

不过，对于简单的字符串处理，你可以使用Lua的内置字符串函数，如string.find来查找"Bearer "的位置，然后使用string.sub来提取token。以下是修改后的代码示例：

```lua
location /test {
    default_type 'text/plain';
    content_by_lua_block {
        -- 获取请求对象
        local req = ngx.req

        -- 获取Authorization请求头字段
        local authorization_header = req.get_header("Authorization")

        -- 检查Authorization请求头是否存在并包含Bearer
        if authorization_header and authorization_header:match("^Bearer ") then
            -- Lua 5.1及更早版本中没有:match的直接支持，这里仅为示例
            -- 实际上，你应该使用string.find和string.sub来提取token
            local pos = string.find(authorization_header, "Bearer ")
            if pos then
                local token = string.sub(authorization_header, pos + 7) -- "Bearer "长度为7
                ngx.say("Bearer Token: ", token)
            else
                ngx.say("Failed to extract Bearer token.")
            end
        else
            ngx.say("Authorization header is not present or does not start with 'Bearer '.")
        end
    }
}
```
请注意，由于Lua 5.1（OpenResty通常使用的版本）的字符串库相对简单，没有内置的:match方法（这是Lua 5.3中引入的），因此上面的:match调用仅作为说明。在实际应用中，你应该使用string.find来查找"Bearer "的位置，并使用string.sub来提取token。

另外，如果你的OpenResty环境支持LuaJIT或你正在使用更新的Lua版本，并且想要使用更现代的字符串处理功能，你可以考虑使用第三方库或升级到支持这些功能的Lua版本。但是，对于大多数基本的字符串处理任务来说，Lua的内置函数已经足够了。

