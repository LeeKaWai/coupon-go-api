package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Config 应用程序配置结构
type Config struct {
	Taobao Taobao `mapstructure:"taobao" json:"taobao" yaml:"taobao"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}

type System struct {
	Env  string `mapstructure:"env" json:"env" yaml:"env"`
	Port string `mapstructure:"port" json:"port" yaml:"port"`
}

type Mysql struct {
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Name     string `mapstructure:"name" json:"name" yaml:"name"`
}

type Taobao struct {
	Appkey           string `mapstructure:"appkey" json:"appkey" yaml:"appkey"`
	Appsecret        string `mapstructure:"appsecret" json:"appsecret" yaml:"appsecret"`
	RestUrl          string `mapstructure:"rest_url" json:"rest_url" yaml:"restUrl"`
	Pid              string `mapstructure:"pid" json:"pid" yaml:"pid"`
	Siteid           string `mapstructure:"siteid" json:"siteid" yaml:"siteid"`
	Adzoneid         string `mapstructure:"adzoneid" json:"adzoneid" yaml:"adzoneid"`
	WechatPid        string `mapstructure:"wechat_pid" json:"wechat_pid" yaml:"wechat_pid"`
	WechatSiteid     string `mapstructure:"wechat_siteid" json:"wechat_siteid" yaml:"wechat_siteid"`
	WechatAdzoneid   string `mapstructure:"wechat_adzoneid" json:"wechat_adzoneid" yaml:"wechat_adzoneid"`
	RelationPid      string `mapstructure:"relation_pid" json:"relation_pid" yaml:"relation_pid"`
	RelationSiteid   string `mapstructure:"relation_siteid" json:"relation_siteid" yaml:"relation_siteid"`
	RelationAdzoneid string `mapstructure:"relation_adzoneid" json:"relation_adzoneid" yaml:"relation_adzoneid"`
	Relationid       string `mapstructure:"relation_id" json:"relation_id" yaml:"relation_id"`
	ZtkAppkey        string `mapstructure:"ztk_appkey" json:"ztk_appkey" yaml:"ztk_appkey"`
	ZtkSid           string `mapstructure:"ztk_sid" json:"ztk_sid" yaml:"ztk_sid"`
}

// 定义全局变量
var GLOBAL_CONF *Config

// LoadConfig 从配置文件加载配置
func LoadConfig(env string) {
	var config Config

	viper := viper.New()

	// 设置配置文件名为当前环境的配置文件（例如：dev.yaml）
	viper.SetConfigName(env)
	viper.AddConfigPath("config")
	viper.SetConfigType("yaml")

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("无法读取配置文件: %s", err)
	}

	// 将配置文件内容映射到结构体
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
	}
	// 设置环境变量
	config.System.Env = env
	// 将配置结构体赋值给全局变量
	GLOBAL_CONF = &config
}
