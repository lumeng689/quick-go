# PPROF 教程

## 安装

```bash

# 安装pprof
go get -u github.com/google/pprof

# 安装graphviz
brew install graphviz

# 生成性能图
go tool pprof http://localhost:8081/debug/pprof/profile

# 采样，默认采集30s的
go tool pprof http://127.0.0.1:8081/debug/pprof/profile -seconds 10

go tool pprof -http=:8082 ~/pprof/pprof.samples.cpu.002.pb.gz

# 查看heap情况
go tool pprof http://127.0.0.1:8081/debug/pprof/heap -seconds 10

# 获取trace
wget http://127.0.0.1:8081/debug/pprof/trace

#查看trace情况
go tool trace -http=127.0.0.1:8000 trace
```


