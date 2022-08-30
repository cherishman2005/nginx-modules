#!/bin/bash
#需要处理的数据
name={"code":200,"success":true,"result":"86f8069971e87fa1"}
echo "处理前的数据"
echo $name
echo "awk对数据进行处理"
#使用多次分割方式获取到需要的数据，并将结果赋值给变量
vars=`echo $name | awk -F':' '{print $NF}'|awk -F "}" '{print $1}'`
#输出awk处理后的数据
echo $vars
