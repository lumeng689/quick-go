package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"quick-go/internal/gin-app/model/sys"
	"strings"
)

var db *gorm.DB

func SetUp() {
	user := "root"
	password := "123456"
	host := "localhost:3306"
	dbName := "test"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbName)

	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	var err error
	if db, err = gorm.Open(mysql.New(mysqlConfig), gormConfig()); err != nil {
		fmt.Printf("create gorm failed::%v \n", err)
	}
}

func initTables() {
	err := db.AutoMigrate(
		sys.User{},
	)

	if err != nil {
		fmt.Printf("create tables failed!! %v\n", err)
	}
}

func gormConfig() *gorm.Config {
	config := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",                              // table name prefix, table for `User` would be `t_users`
			SingularTable: true,                              // use singular table name, table for `User` would be `user` with this option enabled
			NameReplacer:  strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
		},
	}

	return config
}
