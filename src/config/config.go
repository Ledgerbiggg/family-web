package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

var config *GConfig

// GConfig 配置
type GConfig struct {
	Mode        string `yaml:"mode"`
	LogLevel    int    `yaml:"logLevel"`
	ServiceName string `yaml:"serviceName"`
	Address     struct {
		Domain string `yaml:"domain"`
		Ip     string `yaml:"ip"`
		Port   int    `yaml:"port"`
	} `yaml:"address"`
	ServerLevel string `yaml:"serverLevel"`
	Redis       struct {
		Address  string `yaml:"address"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	} `yaml:"redis"`
	Mysql struct {
		Address  string `yaml:"address"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"mysql"`
	Jwt struct {
		SecretKey  string `yaml:"secretKey"`
		ExpireTime int    `yaml:"expireTime"`
	} `yaml:"jwt"`
	Static struct {
		Dir string `yaml:"dir"`
	} `yaml:"static"`
}

func LoadConfig() *GConfig {
	vc := viper.New()

	// 预加载环境变量
	vc.AutomaticEnv()

	// 设置环境变量的分隔符，将点号和横杠替换为下划线
	vc.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))

	// 获取环境变量，判断是加载config还是config-dev
	configFile := "config" // 默认加载 "config"
	if os.Getenv("dev") == "测试分类" {
		configFile = "config-dev" // 如果环境变量dev=测试分类，则加载 "config-dev"
	}

	// 设置要加载的配置文件名称
	vc.SetConfigName(configFile)

	// 添加配置文件路径（当前目录）
	vc.AddConfigPath(".")

	// 设置配置文件类型为 YAML
	vc.SetConfigType("yaml")

	// 读取配置文件
	err := vc.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件出错: %s", err))
	}

	// 将配置文件内容解析到config结构体
	if err := vc.Unmarshal(&config); err != nil {
		log.Panicln("解析配置文件失败: " + err.Error())
	}

	// 返回配置
	return config
}
