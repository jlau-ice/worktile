package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestMysql(t *testing.T) {
	dsn := "root:5626%%..@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn, // 使用上面定义的 DSN 变量
		// 其他配置项保持不变
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	// **关键修改点：检查并处理错误**
	if err != nil {
		t.Fatalf("Failed to connect to MySQL: %v", err)
	}
	fmt.Println(db)
	fmt.Println("Successfully connected to MySQL!")
}
