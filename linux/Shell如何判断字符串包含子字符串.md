# Shell如何判断字符串包含子字符串

## 包含子字符串

```
#!/bin/bash
#
string='hello world'
sub='hello'

if [[ $string =~ $sub ]]
# if [[ $string = *$sub* ]]
# if [[ $string =~ ^.*$sub.*$ ]] # 正则表达式
then
    echo '包含'
else 
    echo '不包含'
fi
```

## 以某个字符串作为开始

```
#!/bin/bash
#
string='hello world'
sub='hello'

if [[ $string = $sub* ]]
#if [[ $string =~ ^$sub.*$ ]] # 正则表达式
then
    echo 'YES'
else 
    echo 'NO'
fi
```

## 以某个字符串作为结束

```
#!/bin/bash
#
string='hello world'
sub='world'

if [[ $string = *$sub ]]
#if [[ $string =~ ^.*$sub$ ]] # 正则表达式
then
    echo 'YES'
else 
    echo 'NO'
fi
```
