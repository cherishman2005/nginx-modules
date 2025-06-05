# Ubuntu：安装rust

## 1.安装curl
```
sudo apt install curl
```
安装速度慢的话可以参考https://blog.csdn.net/cacique111/article/details/125952535更换源


## 2.安装rust
```
curl https://sh.rustup.rs -sSf | sh
```
​
### 2.1选择1)默认安装

### 2.2安装成功

## 3.环境变量生效
```
source $HOME/.cargo/env
```

## 4.查看版本

```
rustc -V
```
