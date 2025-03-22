# prometheus

## api接口

```txt
/api/v1/status/buildinfo

{
  "status": "success",
  "data": {
    "version": "2.15.2+ds",
    "revision": "2.15.2+ds-2",
    "branch": "debian/sid",
    "buildUser": "pkg-go-maintainers@lists.alioth.debian.org",
    "buildDate": "20200113-10:10:54",
    "goVersion": "go1.13.7"
  }
}
```

## prometheus启动

* 启动命令
```
prometheus --config.file="prometheus.yml"
```

### 启动日志

```txt
level=info ts=2025-03-22T13:47:01.967Z caller=head.go:584 component=tsdb msg="replaying WAL, this may take awhile"
level=info ts=2025-03-22T13:47:01.976Z caller=head.go:608 component=tsdb msg="WAL checkpoint loaded"
level=info ts=2025-03-22T13:47:02.074Z caller=head.go:632 component=tsdb msg="WAL segment loaded" segment=50 maxSegment=54
level=info ts=2025-03-22T13:47:02.216Z caller=head.go:632 component=tsdb msg="WAL segment loaded" segment=51 maxSegment=54
level=info ts=2025-03-22T13:47:02.233Z caller=head.go:632 component=tsdb msg="WAL segment loaded" segment=52 maxSegment=54
level=info ts=2025-03-22T13:47:02.285Z caller=head.go:632 component=tsdb msg="WAL segment loaded" segment=53 maxSegment=54
level=info ts=2025-03-22T13:47:02.286Z caller=head.go:632 component=tsdb msg="WAL segment loaded" segment=54 maxSegment=54
level=info ts=2025-03-22T13:47:02.289Z caller=main.go:661 fs_type=EXT4_SUPER_MAGIC
level=info ts=2025-03-22T13:47:02.289Z caller=main.go:662 msg="TSDB started"
level=info ts=2025-03-22T13:47:02.289Z caller=main.go:732 msg="Loading configuration file" filename=prometheus.yml
level=info ts=2025-03-22T13:47:02.292Z caller=main.go:760 msg="Completed loading of configuration file" filename=prometheus.yml
level=info ts=2025-03-22T13:47:02.293Z caller=main.go:615 msg="Server is ready to receive web requests."
```
