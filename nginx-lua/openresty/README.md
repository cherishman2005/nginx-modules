# openresty(nginx-lua)

## json pretty print

- [https://github.com/bungle/lua-resty-prettycjson](https://github.com/bungle/lua-resty-prettycjson)

## lua-resty-jwt

- [https://github.com/SkyLothar/lua-resty-jwt](https://github.com/SkyLothar/lua-resty-jwt)

- [https://opm.openresty.org/package/taylorking/lua-resty-jwt/](https://opm.openresty.org/package/taylorking/lua-resty-jwt/)


* 配置示例
```
        location = /verify {
            content_by_lua_block {
                local cjson = require "cjson"
                local jwt = require "resty.jwt"

                local jwt_token = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9" ..
                    ".eyJmb28iOiJiYXIifQ" ..
                    ".VAoRL1IU0nOguxURF2ZcKR0SGKE1gCbqwyh8u2MLAyY"
                local jwt_obj = jwt:verify("lua-resty-jwt", jwt_token)
                ngx.say(cjson.encode(jwt_obj))
            };
        }
        location = /sign {
            content_by_lua_block {
                local cjson = require "cjson"
                local jwt = require "resty.jwt"

                local jwt_token = jwt:sign(
                    "lua-resty-jwt",
                    {
                        header={typ="JWT", alg="HS256"},
                        payload={foo="bar"}
                    }
                )
                ngx.say(jwt_token)
            };
        }
```

## lua-resty-hmac

- [https://github.com/jkeys089/lua-resty-hmac](https://github.com/jkeys089/lua-resty-hmac)

## lua-resty-http

* http-client客户端请求

- [https://github.com/ledgetech/lua-resty-http](https://github.com/ledgetech/lua-resty-http)

