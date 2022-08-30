# sed和awk

## sed替换（换行）
```
sed 's/},{/},\n{/g'  if.txt  > if_proc.txt 
```

## awk排序

```
cat if_proc.txt | awk -F '"IpPort":' '{print $2}' |  awk -F ',' '{print $1}'  | sort | uniq -c
```
