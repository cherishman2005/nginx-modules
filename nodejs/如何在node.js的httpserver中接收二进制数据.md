# 如何在node.js的httpserver中接收二进制数据

最近使用了protobuf进行数据交互，发送在node.js接收前端的二进制数据出现了数据错误等问题。后来发现思路上面的问题，在req.on('data',()=>{})中的处理不适当才引发数据错乱

正确接收二进制数据代码
```
const server = http.createServer((req, res) => {
        if(req.method==='OPTIONS'){
            res.setHeader("Access-Control-Allow-Origin", "*");
            res.statusCode=200;
        }
        if(req.method==='POST'){
          // 存储数组空间
          let msg=[];
          // 接收到数据消息
          req.on('data',(chunk)=>{
            if(chunk){
              msg.push(chunk);
            }
          })
          // 接收完毕
          req.on('end',()=>{
            // 对buffer数组阵列列表进行buffer合并返回一个Buffer
            let buf=Buffer.concat(msg);
            conosole.log(buf)//提取Buffer正确
          })                  
}    
});
server.listen(3000,'127.0.0.1');
```
通过上面可以看到我用数据的方式来存储req发送过来的buffer实例。然后通过Buffer.concat的方式将buffer数组阵列的数据进行合并成一个buffer就行了

这样就可以获取完成的二进制数据了。
