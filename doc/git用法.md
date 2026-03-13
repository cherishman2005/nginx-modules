# git用法

## git显示作者和时间

```
git log --oneline -n 1000 --pretty=format:"%h %an %ad %s"
```

## blame精确定位提交记录

* 精确到文件、行
```
git blame -L 2000,2040  mod.rs 
```
