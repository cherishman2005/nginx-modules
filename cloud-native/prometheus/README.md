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
level=info ts=2025-03-22T13:47:01.950Z caller=repair.go:59 component=tsdb msg="found healthy block" mint=1742472000000 maxt=1742536800000 ulid=01JPVYWHSGK1H7WRDY1PKKJ0YN
level=info ts=2025-03-22T13:47:01.950Z caller=repair.go:59 component=tsdb msg="found healthy block" mint=1742536800000 maxt=1742601600000 ulid=01JPXWP34VQDYRZC3E4481YSJ0
level=info ts=2025-03-22T13:47:01.950Z caller=repair.go:59 component=tsdb msg="found healthy block" mint=1742623200000 maxt=1742630400000 ulid=01JPYH98HBHHP5N7JN1MT2YBVC
level=info ts=2025-03-22T13:47:01.950Z caller=repair.go:59 component=tsdb msg="found healthy block" mint=1742601600000 maxt=1742623200000 ulid=01JPYH98MP9XWJAWAYKF9SS89P
level=info ts=2025-03-22T13:47:01.950Z caller=repair.go:59 component=tsdb msg="found healthy block" mint=1742630400000 maxt=1742637600000 ulid=01JPYR4ZSB43YGSF9228XTWDMS
level=info ts=2025-03-22T13:47:01.950Z caller=repair.go:59 component=tsdb msg="found healthy block" mint=1742637600000 maxt=1742644800000 ulid=01JPYZ0P50BQB5M9VZ7CZE93N3
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
