# bypass_AV

## go_adduser.go

调用Windows API 添加用户的实现，可bypass 360、Windows Denfender、其他杀软自行测试。

编译方法：

```bash
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build go_adduer.go
```
