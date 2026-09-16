**Ubuntu 22.04 安装 OpenResty（官方 APT 仓库方式，推荐）**

官方文档针对 Ubuntu 22 及以上版本使用新的 GPG 密钥管理方式。

### 1. 如果已安装 Nginx，先停止并禁用（避免冲突）
```bash
sudo systemctl stop nginx
sudo systemctl disable nginx
```

### 2. 安装依赖
```bash
sudo apt-get -y install --no-install-recommends wget gnupg ca-certificates lsb-release
```

### 3. 导入 OpenResty GPG 密钥（Ubuntu 22+ 专用）
```bash
wget -O - https://openresty.org/package/pubkey.gpg | sudo gpg --dearmor -o /usr/share/keyrings/openresty.gpg
```

### 4. 添加官方 APT 仓库

**x86_64 / amd64 系统：**
```bash
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/openresty.gpg] https://openresty.org/package/ubuntu $(lsb_release -sc) main" \
    | sudo tee /etc/apt/sources.list.d/openresty.list > /dev/null
```

**ARM64 / aarch64 系统：**
```bash
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/openresty.gpg] https://openresty.org/package/arm64/ubuntu $(lsb_release -sc) main" \
    | sudo tee /etc/apt/sources.list.d/openresty.list > /dev/null
```

### 5. 更新并安装
```bash
sudo apt-get update
sudo apt-get -y install openresty
```

（会顺带推荐安装 `openresty-opm` 和 `openresty-restydoc`。如果不想安装推荐包，可加 `--no-install-recommends`）

### 6. 启动并验证
```bash
sudo systemctl start openresty
sudo systemctl enable openresty   # 开机自启（可选）
sudo systemctl status openresty

# 检查版本
openresty -V

# 测试默认页面
curl 127.0.0.1/
```

成功后会看到 OpenResty 的欢迎页面。

### 常用路径
- 默认安装前缀：`/usr/local/openresty/`
- 可执行文件：`/usr/bin/openresty`（软链到 nginx）
- 配置文件：`/usr/local/openresty/nginx/conf/nginx.conf`
- 日志：`/usr/local/openresty/nginx/logs/`

### 其他常用命令
```bash
sudo systemctl reload openresty   # 重载配置
sudo systemctl restart openresty  # 重启
```
