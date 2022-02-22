#!/bin/bash
######################################################################
#
#实现ab多url并发的shell脚本
#
# sh ab.sh 并发请求数(-c) 最大秒数(-t)  请求的次数(-n)
#
######################################################################
rm -rf ab.log  #清空日志
for i in $(cat url.txt)
do
    if [ "$1" == "" ]
    then 
        echo "并发请求数不能为空" 
    elif [[ "$2" == "" ]]
    then
        if [ "$3" == "" ]
        then
            echo "并发请求数 = $1,最大秒数未赋值，请求的次数未赋值"
            ab -c $1 $i >> ab.log &
            continue
        else
            echo "并发请求数 = $1,最大秒数未赋值，请求的次数 = $3 "
            ab -t $2 -n $3 $i >> ab.log &
            continue
        fi

    elif [[ "$3" == "" ]]
    then
        echo "并发请求数 = $1,最大秒数 = $2 ,请求的次数未赋值"
        ab -c $1 -t $2 $i >> ab.log &
        continue
    else
        echo "并发请求数 = $1,最大秒数 = $2 ,请求的次数 = $3 "
        ab  -c $1 -t $2 -n $3 $i >> ab.log &
        continue
    fi
done

#####################################################################
#
# for 循环读取url文件中的url内容，执行ab命令
#
# 接收的url是i变量(从URL中读取的每行url值)
#
# 将结果写入ab.log日志中，& shell中是并行
#
# 1 是 sh ab.sh x xx xxx第一个x的值表示并发请求数
#
# 2 是 sh ab.sh x xx xxx第二个xx的值表示测试所进行的最大秒数
#
# 3 是 sh ab.sh x xx xxx第三个xxx的值表示每次ab请求的次数
#
####################################################################