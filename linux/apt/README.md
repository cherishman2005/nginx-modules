# apt-get

This is solution fix this issue on ubuntu server 14.04.x

1, Edit file:
```
sudo vi  /etc/apt/sources.list
```

2, Add to file sources.list
```
deb http://security.ubuntu.com/ubuntu xenial-security main
deb http://cz.archive.ubuntu.com/ubuntu xenial main universe
```

3, Run command update and update CURL to new version
```
apt-get update && apt-get install curl
```

- [https://stackoverflow.com/questions/60262230/fatal-unable-to-access-gnutls-handshake-failed-handshake-failed](https://stackoverflow.com/questions/60262230/fatal-unable-to-access-gnutls-handshake-failed-handshake-failed)


# FAQ

## 报错分析
```
W: The repository 'https://storage.googleapis.com/bazel-apt stable Release' does not have a Release file.
N: Data from such a repository can't be authenticated and is therefore potentially dangerous to use.
N: See apt-secure(8) manpage for repository creation and user configuration details.
W: The repository 'https://mirrors.aliyun.com/kubernetes/apt kubernetes-xenial Release' does not have a Release file.
N: Data from such a repository can't be authenticated and is therefore potentially dangerous to use.
N: See apt-secure(8) manpage for repository creation and user configuration details.
E: Failed to fetch https://storage.googleapis.com/bazel-apt/dists/stable/jdk1.8/binary-amd64/Packages  gnutls_handshake() failed: Decryption has failed.
E: Failed to fetch https://mirrors.aliyun.com/kubernetes/apt/dists/kubernetes-xenial/main/binary-amd64/Packages  server certificate verification failed. CAfile: /etc/ssl/certs/ca-certificates.crt CRLfile: none
E: Some index files failed to download. They have been ignored, or old ones used instead.
```

sources.list里面没有这2个文件，不知为啥多出来：

```
-rw-r--r-- 1 root root   72 Jan 12 19:35 bazel.list
-rw-r--r-- 1 root root   72 Jan 12 19:35 bazel.list.save
-rw-r--r-- 1 root root   70 Jan 12 19:35 kubernetes.list
-rw-r--r-- 1 root root   70 Jan 12 19:35 kubernetes.list.save
```
删除后 apt-get update ok
