# Linux下去掉^M的方法


在Linux下使用vi来查看一些在Windows下创建的文本文件，有时会发现在行尾有一些“^M”。有几种方法可以处理。


方法一：

使用cat命令就可以看到windows下的断元字符 ^M
```
cat -A filename 
```
要去除他，最简单用下面的命令：
```
dos2unix filename
```

方法二：
```
sed -i 's/^M//g' filename
        #注意：^M的输入方式是 Ctrl + v ，然后Ctrl + M
```

方法三：

使用vi的替换功能。启动vi，进入命令模式，输入以下命令：
```
:%s/^M$//g # 去掉行尾的^M。
 
:%s/^M//g # 去掉所有的^M。
 
:%s/^M/[ctrl-v]+[enter]/g # 将^M替换成回车。
 
:%s/^M/\r/g # 将^M替换成回车。
        #注意：^M 输入方法： ctrl+V ,ctrl+M
```

方法四：
```
cat filename |tr -d '/r' > newfile
        #^M 可用 /r 代替
```
