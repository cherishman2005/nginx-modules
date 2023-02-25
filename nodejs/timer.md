# timer


## 定时器

```
let count = 5;
let i = 0;
const intervalObj = setInterval((arg) => {
    console.log(`arg was => ${arg}`);
    if (++i >= count) {
        clearInterval(intervalObj);
        console.log("stop")
    }
}, 1500, 'funky');
```


# 参考链接

- [Node.js 中的定时器](https://nodejs.org/zh-cn/docs/guides/timers-in-node/)
