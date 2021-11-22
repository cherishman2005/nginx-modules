local cjson = require "cjson"
-- read data --
ngx.req.read_body()
local data = ngx.req.get_body_data()
if not data then
    return ngx.exit(403)
end
ngx.log(ngx.INFO, 'auth_data="' .. data .. '"')

local obj = cjson.decode(data);
if not obj or not obj.username or not obj.token or not obj.bos_path then
    ngx.log(ngx.ERR, "obj params is nil")
    return ngx.exit(403)
end

local out = { username = obj.username, token = obj.token, bos_path = obj.bos_path, code = 200, msg = "ok" }
ngx.print(cjson.encode(out));