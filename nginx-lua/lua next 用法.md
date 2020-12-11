(转) 关于lua table是否为空的判断
在项目的脚本lua中经常有这样的需求，

1、local a = {}

2、对a进行处理

3、对a是否为空表进行判断

关于对a是否为空表的判断，我发现有些代码如此做：

if a == {} then

这样的结果就是a == {}永远返回false，是一个逻辑错误。因为这里比较的是table a和一个匿名table的内存地址。

也有些代码如此做：

if table.maxn(a) == 0 then

这样做也不保险，除非table的key都是数字，而没有hash部分。

难道真的要遍历table发现有东西就return false跳出才能断定它是否为空吗？这样写至少代码太难看. 

网上小搜了一下，发现原来官方手册里早已经给了答案，那就是`靠lua内置的next函数`

即如此用：if next(a) == nil then

next其实就是pairs遍历table时用来取下一个内容的函数。

在项目的module中最好封装一下，免得module本地也有next函数

于是封装后判断的lua table是否为空的函数如下：

function table_is_empty(t)
        return _G.next( t ) == nil
end


------------------------------------------------

自己也用了一下local playerId = next(MsgGuildData.getGuildMemberList())

OK很好用

这样就不用循环去遍历取key表的第一个值了