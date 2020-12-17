package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// MySQLConfig 数据库配置
type MySQLConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DB       string `json:"db"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

// RedisConfig redis 配置
type RedisConfig struct {
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

// Config 项目配置
type Config struct {
	Port  int         `json:"port"`
	Name  string      `json:"name,omitempty"`
	Mysql MySQLConfig `json:"mysql,omitempty"`
	Redis RedisConfig `json:"redis,omitempty"`
}

// ReadConfig 读取配置
func ReadConfig() Config {
	var runtimeViper = viper.New()
	runtimeViper.SetConfigFile("config.yaml") // 指定配置文件
	runtimeViper.AddConfigPath("../config/")  // 多次调用以添加多个搜索路径
	runtimeViper.AddConfigPath(".")           // 还可以在工作目录中查找配置
	runtimeViper.AddConfigPath("../")

	var runtimeConf Config

	// read from config the first time.
	err := runtimeViper.ReadInConfig()
	if err != nil { // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// unmarshal config
	runtimeViper.Unmarshal(&runtimeConf)

	// open a goroutine to watch changes forever
	go func() {
		for {
			time.Sleep(time.Second * 5) // delay after each request

			// currently, only tested with etcd support
			runtimeViper.WatchConfig()

			// unmarshal new config into our runtime config struct. you can also use channel
			// to implement a signal to notify the system of the changes
			runtimeViper.Unmarshal(&runtimeConf)
		}
	}()
	return runtimeConf
}
