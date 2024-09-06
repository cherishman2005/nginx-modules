local table = {}

function table.kIn(tbl, key)
    if tbl == nil then
        return false
    end
    for k, v in pairs(tbl) do
        if v == key then
            return true
        end
    end
    return false
end

local tbl = {"127.0.0.1", "127.0.0.2"}
local key = "127.0.0.2"
print(table.kIn(tbl, key))
