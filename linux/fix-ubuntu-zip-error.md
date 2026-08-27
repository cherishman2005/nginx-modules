# Ubuntu 22.04 修复 zip: No such file or directory (os error 2)

## 安装 zip / unzip

```bash
sudo apt update
sudo apt install zip unzip -y
```

## 校验安装

```bash
which zip
zip -v
```

正常输出示例：

```
/usr/bin/zip
This is Zip 3.0 (July 5th 2008), by Info-ZIP
```

---

## 常见问题

### 1. 装完依旧报错

- 程序运行环境是 docker 容器：**必须在容器内部安装**，宿主机安装无效。进入容器再执行上面 apt 命令，或者写到 Dockerfile：

```dockerfile
RUN apt update && apt install -y zip unzip
```

### 2. 程序使用绝对路径调用 zip

拿到 `which zip` 的输出路径，在代码里填完整路径 `/usr/bin/zip`，不要只写 `zip`。

### 3. 权限问题

os error 2 不是权限，是找不到程序；权限不足是 os error 13。

---

## 快速测试

```bash
zip test.zip test.txt
```

不报错就代表系统 zip 工具正常可用。
