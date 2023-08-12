# 获取主机信息
---


# 编译
## windows

64 位
```
GOOS=windows GOARCH=amd64 go build -o sysinfo.exe main.go
```

32 位
```
GOOS=windows GOARCH=386 go build -o sysinfo_32.exe main.go
```

## linux
```
GOOS=linux GOARCH=amd64 go build -o sysinfo main.go
```

## darwin
```
GOOS=darwin GOARCH=amd64 go build -o sysinfo main.go
```