# Ubuntu date 显示 UTC 时间完整修复方案

> 适用场景：`date` 输出为 UTC 时间，需要改为中国标准时间 CST（UTC+8）

---

## 一、先查看当前时间配置

执行命令查看全部时间信息：

```bash
timedatectl
```

重点观察以下字段：

| 字段 | 含义 |
|------|------|
| `Time zone` | 时区设置，`Etc/UTC` 表示时区错误 |
| `RTC in local TZ` | 硬件时钟是否使用本地时间 |
| `System clock synchronized` | NTP 网络同步状态 |

---

## 二、标准修复（单系统 Ubuntu，推荐）

### 1. 修改时区为中国东八区

```bash
sudo timedatectl set-timezone Asia/Shanghai
```

> 执行完立刻生效，无需重启。再次执行 `date` 验证，正常会显示 `CST`（中国标准时间，UTC+8）。

### 2. 开启网络自动对时（防止时间跑偏）

```bash
# 开启 NTP 自动同步
sudo timedatectl set-ntp true

# 重启时间同步服务
sudo systemctl restart systemd-timesyncd
```

### 3. 同步硬件 RTC 时钟（关机重启时间不变）

Linux 标准规范：硬件时钟 RTC 使用 UTC，**不要改成本地时间**。

```bash
# RTC 使用 UTC（最佳实践）
sudo timedatectl set-local-rtc 0

# 把当前正确系统时间写入主板硬件时钟
sudo hwclock --systohc
```

### 验证

```bash
date
timedatectl
```

---

## 三、Windows + Ubuntu 双系统时间错乱专用修复

**问题根源**：Windows 默认把硬件时钟当本地时间，Linux 默认把硬件时钟当 UTC，互相冲突、来回差 8 小时。

### 方案 A：让 Ubuntu 适配 Windows（简单常用）

让 Ubuntu 读取硬件时钟为本地时间：

```bash
sudo timedatectl set-timezone Asia/Shanghai
sudo timedatectl set-local-rtc 1 --adjust-system-clock
sudo hwclock --systohc
```

### 方案 B：改 Windows 适配 Linux（推荐长期稳定）

Windows 注册表强制硬件时钟使用 UTC，不用改 Linux：

在 Windows 中以**管理员**打开命令提示符，执行：

```cmd
reg add HKLM\SYSTEM\CurrentControlSet\Control\TimeZoneInformation /v RealTimeIsUniversal /t REG_DWORD /d 1 /f
```

Windows 重启后，Linux 保持默认 `set-local-rtc 0` 即可。

---

## 四、timedatectl 无效时，手动硬改时区（容器 / 极简系统）

```bash
# 删除旧时区软链接
sudo rm -f /etc/localtime

# 软链接到上海时区文件
sudo ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# 写入时区配置文件
echo "Asia/Shanghai" | sudo tee /etc/timezone
```

### 交互式配置时区

```bash
sudo dpkg-reconfigure tzdata
```

弹窗依次选择 **Asia → Shanghai**。

---

## 五、NTP 网络同步失败、时间不准的进阶修复

国内 NTP 换成阿里云 / 腾讯服务器，解决外网 NTP 超时问题。

### 创建自定义 NTP 配置

```bash
sudo mkdir -p /etc/systemd/timesyncd.conf.d
sudo nano /etc/systemd/timesyncd.conf.d/cn-ntp.conf
```

写入以下内容：

```ini
[Time]
NTP=ntp.aliyun.com ntp.tencent.com cn.pool.ntp.org
FallbackNTP=ntp.ubuntu.com
```

保存退出：`Ctrl+O` 回车保存，`Ctrl+X` 退出。

### 重载服务并强制同步

```bash
sudo systemctl restart systemd-timesyncd
sudo timedatectl set-ntp true
```

### 手动强制拉取网络时间（需要先临时关闭 NTP）

```bash
sudo timedatectl set-ntp false
sudo apt install ntpdate -y
sudo ntpdate ntp.aliyun.com
sudo timedatectl set-ntp true
```

---

## 六、常见问题排查

### 改完时区部分程序（日志、Java、Docker）依旧 UTC

程序启动时读取时区，需重启对应服务：

```bash
sudo systemctl restart cron rsyslog docker
```

Java 可临时加环境变量：

```bash
export TZ=Asia/Shanghai
```

### 重启 Ubuntu 时间又变回 UTC

检查是否装了 `chrony` / `ntpd` 多个时间服务冲突：

```bash
sudo systemctl stop chronyd ntpd
sudo systemctl disable chronyd ntpd
```

确认执行过 `sudo hwclock --systohc` 写入硬件时钟。

### 时区文件缺失

重装时区数据包：

```bash
sudo apt update && sudo apt install --reinstall tzdata
```

---

## 最终校验命令

```bash
date
timedatectl
hwclock --show
```

### 正常 `date` 输出示例

```
Sun Aug  2 16:04:08 CST 2026
```

> 比 UTC 快 8 小时，即为修复成功。
