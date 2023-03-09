# linux shell报错Syntax error: Bad for loop variable

在linux下写了一个简单的shell，循环10次．
test.sh
```
#!/bin/bash                        
                            
##循环10次                        
for ((i=0; i<10; i++));                                         
do                                                     
    echo $i                                                                                                     
done
```
执行：sh test.sh 报下面的错误．
``
Syntax error: Bad for loop variable
```

因为Ubuntu为了加快开机速度，用dash代替了传统的bash，所以我们这样执行就没问题．
``
bash test.sh
```
 

那如果我们只想用sh test.sh 这样的方式执行，怎么办呢？

修改一下代码．
```
for i in `seq 10`                  
do                                 
                                   
    echo Good Morning ,this is  $i  shell program.
                                   
done
```
这个时候，你再执行 sh test.sh，就不会报错误啦．
