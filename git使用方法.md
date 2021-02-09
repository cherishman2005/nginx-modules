# git将一个分支完全覆盖另外一个分支如：dev-1.0分支代码完全覆盖master分支

如：当前分支是maser分支，我想讲dev-1.0分支上的代码完全覆盖master分支，首先切换到master分支。
```
git reset --hard origin/dev-1.0
```

执行上面的命令后master分支上的代码就完全被dev-1.0分支上的代码覆盖了（本地分支），然后将本地分支强行推到远程分支。
```
git push -f
```

