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

```
https://github.com/clash-verge-rev/clash-verge-rev/releases/download/v2.4.4/Clash.Verge_2.4.4_amd64.deb
```

# http_proxy代理设置

## 设置系统代理（临时生效）

```bash
# 设置HTTP代理
export http_proxy=http://127.0.0.1:7890
export https_proxy=http://127.0.0.1:7890

# 取消代理
unset http_proxy https_proxy
```
