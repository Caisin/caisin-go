# caisin-go

## setup
```shell
go get github.com/Caisin/caisin-go
```

## 设置私有库
```
# 告诉 Go 工具链不通过代理拉取私有仓库,否则无法拉取到私有仓库
go env -w GOPRIVATE="gitee.com/Caisin,github.com/Caisin"
```