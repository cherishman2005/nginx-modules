# 构建hello world Docker镜像



简介： Docker镜像构建是通过Dockerfile来构建的，里面运行的程序是可以自定的，从编写程序到安装Docker镜像，可以一气呵成。接下来我们就通过九步实现一个自定义的镜像的制作、构建及运行。

## 构建 Docker 镜像

Docker镜像构建是通过 Dockerfile来构建的，里面运行的程序是可以自定的，从编写程序到安装Docker镜像，可以一气呵成。接下来我们就通过九步实现一个自定义的镜像的制作、构建及运行。

编写在Docker中要运行的程序

该部分主要是镜像中程序的编写及编译

### 1. 创建目录
```
# 创建应用程序目录
mkdir hello-docker
# 进入该目录
cd hello-docker
```

### 2. 编写程序
```
vi hello.cpp
```

### 3. 源码

```
#include <iostream>
using namespace std;

int main() {
    cout << "hello docker" << endl;
    return 0;
}
```


### 4. 编译

将编写的cpp程序编译在hello-docker根目录
```
g++ hello.cpp -o hello
```

Dockerfile文件编写

### 5. 在hello-docker文件夹下新建Dockerfile文件
```
vim Dockerfile
```

### 6. 编写Dockerfile中的内容
```
FROM scatch
ADD hello /
RUN apt-get update
CMD ["/hello"]
```

注： FROM:从哪构建镜像，是基础镜像的地址或者名称;ADD: 添加文件到镜像执行位置，如上是根目录;RUN: 镜像的操作指令;CMD: 容器启动时执行指令

编译及运行

### 7. 在hello-docker 根目录编译镜像

```
docker build -t zhangbiwu/test .
```

### 8. 编译完成后，在命令行中查看镜像是否存在

```
docker images
```

### 9. 运行镜像

```
docker run guzhongren/test
```

总结

知其然，知其所以然。
