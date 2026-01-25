# linux_clash配置

## 创建Clash配置目录

```
mkdir -p ~/.config/clash
cd ~/.config/clash
```

## 下载Clash二进制文件（Linux AMD64）

```
wget https://github.com/Dreamacro/clash/releases/download/premium/clash-linux-amd64-2024.12.15.gz
gzip -d clash-linux-amd64-2024.12.15.gz
chmod +x clash-linux-amd64-2024.12.15
mv clash-linux-amd64-2024.12.15 /usr/local/bin/clash
```
