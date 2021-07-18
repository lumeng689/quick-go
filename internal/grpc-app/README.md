# grpc-app

## 开发环境准备

### gRPC准备

访问地址[https://github.com/protocolbuffers/protobuf/releases/](https://github.com/protocolbuffers/protobuf/releases/)，找到对应版本的mac包，下载protoc命令，并安装

```bash
# 安装go生成的插件
go get -u -v github.com/golang/protobuf/protoc-gen-go
```

### 生成protobuf文件

```bash
protoc ./protobuf/*/*.proto --go_out=plugins=grpc:./../../..
```

