module quick-go

go 1.14

require (
	github.com/360EntSecGroup-Skylar/excelize/v2 v2.4.0
	github.com/apache/thrift v0.14.2
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/gin-contrib/pprof v1.3.0
	github.com/gin-gonic/gin v1.7.2
	github.com/go-redis/redis/v8 v8.11.1
	github.com/google/uuid v1.1.2
	github.com/graphql-go/graphql v0.7.9
	github.com/influxdata/influxdb-client-go/v2 v2.4.0
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/json-iterator/go v1.1.11
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.5 // indirect
	github.com/prometheus/client_golang v1.11.0
	github.com/spf13/viper v1.8.1
	go.uber.org/zap v1.18.1
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.26.0
	gorm.io/driver/mysql v1.1.1
	gorm.io/gorm v1.21.12
)

replace github.com/360EntSecGroup-Skylar/excelize/v2 v2.4.1-0.20210728160357-7dbf88f221f2 => github.com/xuri/excelize/v2 v2.4.1-0.20210728160357-7dbf88f221f2
