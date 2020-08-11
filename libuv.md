# libuv原理

libuv 是 Node 的新跨平台抽象层，用于抽象 Windows 的 IOCP 及 Unix 的 libev。

![libuv原理](/img/libuv.png)

特性：
* 非阻塞 TCP 套接字
* 非阻塞命名管道
* UDP
* 定时器
* 子进程生成
* 通过 uv_getaddrinfo 实现异步 DNS
* 异步文件系统 API：uv_fs_*
* 高分辨率时间：uv_hrtime
* 正在运行程序路径查找：uv_exepath
* 线程池调度：uv_queue_work
* TTY控制的ANSI转义代码: uv_tty_t
* 文件系统事件现在支持 inotify, ReadDirectoryChangesW 和 kqueue。很快会支持事件端口：uv_fs_event_t
* 进程间的 IPC 与套接字共享：uv_write2