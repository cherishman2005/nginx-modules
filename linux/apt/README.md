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
