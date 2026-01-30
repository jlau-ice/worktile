package config

import (
	"fmt"
	"net/url"

	"github.com/spf13/viper"
)

// Config 应用配置
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port string
	Mode string // debug, release, test
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbName"`
	TimeZone string `mapstructure:"timeZone"`
}

// LoadConfig 加载配置
// 返回 *Config 以便依赖注入容器使用
func LoadConfig() (*Config, error) {
	v := viper.New()
	v.SetConfigName("config")   // 配置文件名称(不带扩展名)
	v.SetConfigType("yaml")     // 配置文件类型
	v.AddConfigPath("./config") // config子目录
	v.AddConfigPath(".")        // 当前目录

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("加载配置文件失败: %w", err)
	}
	// 支持环境变量覆盖配置（可选）
	v.AutomaticEnv()

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("配置绑定失败: %w", err)
	}
	return &cfg, nil
}

// GetDSN 获取数据库连接字符串
func (c *Config) GetDSN() string {
	escapedPass := url.QueryEscape(c.Database.Password)
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=%s",
		c.Database.User,
		escapedPass,
		c.Database.Host,
		c.Database.Port,
		c.Database.DBName,
	)
}
