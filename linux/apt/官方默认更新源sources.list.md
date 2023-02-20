# 官方默认更新源sources.list

第一步，备份官方的默认源

避免自己手贱操作失误，重装系统太费时间

```
cp /etc/apt/sources.list /etc/apt/sources.list.bak
```
此命令是备份该官方源文件

## Ubuntu 14.04.4官方默认更新源sources.list
```
# deb cdrom:[Ubuntu 14.04.4 LTS _Trusty Tahr_ - Release amd64 (20160217.1)]/ trusty main restricted

# See http://help.ubuntu.com/community/UpgradeNotes for how to upgrade to
# newer versions of the distribution.
deb http://cn.archive.ubuntu.com/ubuntu/ trusty main restricted
deb-src http://cn.archive.ubuntu.com/ubuntu/ trusty main restricted

## Major bug fix updates produced after the final release of the
## distribution.
deb http://cn.archive.ubuntu.com/ubuntu/ trusty-updates main restricted
deb-src http://cn.archive.ubuntu.com/ubuntu/ trusty-updates main restricted

## N.B. software from this repository is ENTIRELY UNSUPPORTED by the Ubuntu
## team. Also, please note that software in universe WILL NOT receive any
## review or updates from the Ubuntu security team.
deb http://cn.archive.ubuntu.com/ubuntu/ trusty universe
deb-src http://cn.archive.ubuntu.com/ubuntu/ trusty universe
deb http://cn.archive.ubuntu.com/ubuntu/ trusty-updates universe
deb-src http://cn.archive.ubuntu.com/ubuntu/ trusty-updates universe

## N.B. software from this repository is ENTIRELY UNSUPPORTED by the Ubuntu
## team, and may not be under a free licence. Please satisfy yourself as to
## your rights to use the software. Also, please note that software in
## multiverse WILL NOT receive any review or updates from the Ubuntu
## security team.
deb http://cn.archive.ubuntu.com/ubuntu/ trusty multiverse
deb-src http://cn.archive.ubuntu.com/ubuntu/ trusty multiverse
deb http://cn.archive.ubuntu.com/ubuntu/ trusty-updates multiverse
deb-src http://cn.archive.ubuntu.com/ubuntu/ trusty-updates multiverse

## N.B. software from this repository may not have been tested as
## extensively as that contained in the main release, although it includes
## newer versions of some applications which may provide useful features.
## Also, please note that software in backports WILL NOT receive any review
## or updates from the Ubuntu security team.
deb http://cn.archive.ubuntu.com/ubuntu/ trusty-backports main restricted universe multiverse
deb-src http://cn.archive.ubuntu.com/ubuntu/ trusty-backports main restricted universe multiverse

deb http://security.ubuntu.com/ubuntu trusty-security main restricted
deb-src http://security.ubuntu.com/ubuntu trusty-security main restricted
deb http://security.ubuntu.com/ubuntu trusty-security universe
deb-src http://security.ubuntu.com/ubuntu trusty-security universe
deb http://security.ubuntu.com/ubuntu trusty-security multiverse
deb-src http://security.ubuntu.com/ubuntu trusty-security multiverse

## Uncomment the following two lines to add software from Canonical's
## 'partner' repository.
## This software is not part of Ubuntu, but is offered by Canonical and the
## respective vendors as a service to Ubuntu users.
# deb http://archive.canonical.com/ubuntu trusty partner
# deb-src http://archive.canonical.com/ubuntu trusty partner

## This software is not part of Ubuntu, but is offered by third-party
## developers who want to ship their latest software.
deb http://extras.ubuntu.com/ubuntu trusty main
deb-src http://extras.ubuntu.com/ubuntu trusty main
```
 

 

## Ubuntu 16.04官方默认更新源sources.list

```
#deb cdrom:[Ubuntu 16.04.1 LTS _Xenial Xerus_ - Release amd64 (20160719)]/ xenial main restricted

# See http://help.ubuntu.com/community/UpgradeNotes for how to upgrade to
# newer versions of the distribution.
deb http://cn.archive.ubuntu.com/ubuntu/ xenial main restricted
# deb-src http://cn.archive.ubuntu.com/ubuntu/ xenial main restricted

## Major bug fix updates produced after the final release of the
## distribution.
deb http://cn.archive.ubuntu.com/ubuntu/ xenial-updates main restricted
# deb-src http://cn.archive.ubuntu.com/ubuntu/ xenial-updates main restricted

## N.B. software from this repository is ENTIRELY UNSUPPORTED by the Ubuntu
## team, and may not be under a free licence. Please satisfy yourself as to
## your rights to use the software. Also, please note that software in
## universe WILL NOT receive any review or updates from the Ubuntu security
## team.
deb http://cn.archive.ubuntu.com/ubuntu/ xenial universe
# deb-src http://cn.archive.ubuntu.com/ubuntu/ xenial universe
deb http://cn.archive.ubuntu.com/ubuntu/ xenial-updates universe
# deb-src http://cn.archive.ubuntu.com/ubuntu/ xenial-updates universe
```
