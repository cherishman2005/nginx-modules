# timersub

```
#define	timersub(a, b, result)						      \
  do {									      \
    (result)->tv_sec = (a)->tv_sec - (b)->tv_sec;			      \
    (result)->tv_usec = (a)->tv_usec - (b)->tv_usec;			      \
    if ((result)->tv_usec < 0) {					      \
      --(result)->tv_sec;						      \
      (result)->tv_usec += 1000000;					      \
    }									      \
  } while (0)
```

# 参考链接

- [https://chromium.googlesource.com/native_client/nacl-newlib/+/99fc6c167467b41466ec90e8260e9c49cbe3d13c/newlib/libc/include/sys/time.h](https://chromium.googlesource.com/native_client/nacl-newlib/+/99fc6c167467b41466ec90e8260e9c49cbe3d13c/newlib/libc/include/sys/time.h)
