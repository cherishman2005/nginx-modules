# quic编译

```
cd build && cmake .. && make -j

cd benchmark/server && cmake . && make 

cd benchmark/client && cmake . && make

GCC >= 7 + CMAKE > 3.10  + GO
```

![75B56E9F0F20B6DA76E4D009F55275C0](https://user-images.githubusercontent.com/17688273/173582758-03651485-fe60-4b48-9480-3d16ba2947f9.jpg)


![528A1B739FE3C154060ADD0C75C0A87D](https://user-images.githubusercontent.com/17688273/173582889-432e3ae9-cadc-4419-bb03-ad1f17f9a137.jpg)


https://github.com/microsoft/mimalloc


# 参考链接

- [https://tools.ietf.org/id/draft-ietf-quic-recovery-27.html](https://tools.ietf.org/id/draft-ietf-quic-recovery-27.html)
