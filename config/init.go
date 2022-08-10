package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
	"sync"
)

type EngineConfig struct {
	Sink   Sink   `yaml:"sink"`
	Source Source `yaml:"source"`
	Map    Map    `yaml:"map"`
	Keyby  Keyby  `yaml:"keyby"`
	Reduce Reduce `yaml:"reduce"`
}

type Keyby struct {
	Lb   string `yaml:"lb"`
	Num  int    `yaml:"num"`
	Ip   string `yaml:"ip"`
	Port int    `yaml:"port"`
}

type Reduce struct {
	Num  int    `yaml:"num"`
	Ip   string `yaml:"ip"`
	Port int    `yaml:"port"`
}

type Sink struct {
	Num  int    `yaml:"num"`
	Ip   string `yaml:"ip"`
	Port int    `yaml:"port"`
}

type Source struct {
	Num  int    `yaml:"num"`
	Ip   string `yaml:"ip"`
	Port int    `yaml:"port"`
	Lb   string `yaml:"lb"`
}

type Map struct {
	Ip   string `yaml:"ip"`
	Port int    `yaml:"port"`
	Lb   string `yaml:"lb"`
	Num  int    `yaml:"num"`
}

var Config *EngineConfig
var once sync.Once
var GlobalDAGConfig *viper.Viper

func init() {
	GlobalDAGConfig = InitDAGConfig()
	dynamicDAGConfig()
}

func InitEngine() *EngineConfig {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("获取文件路径失败：", err)
	}
	c := &EngineConfig{}
	v := viper.New()
	v.SetConfigName("config")       //配置文件名称
	v.AddConfigPath(wd + "/config") //文件所在的目录路径
	v.SetConfigType("yml")          //文件格式类型
	err = v.ReadInConfig()
	if err != nil {
		log.Fatal("读取配置文件失败：", err)
	}
	configs := v.AllSettings()
	for k, val := range configs {
		v.SetDefault(k, val)
	}
	err = v.Unmarshal(c) //反序列化至结构体
	if err != nil {
		log.Fatal("读取配置错误：", err)
	}
	return c
}

func GetConfig() *EngineConfig {
	once.Do(func() {
		Config = InitEngine()
	})
	return Config
}

func InitDAGConfig() *viper.Viper {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("获取文件路径失败：", err)
	}
	v := viper.New()
	v.SetConfigName("DAG")          //配置文件名称
	v.AddConfigPath(wd + "/config") //文件所在的目录路径
	v.SetConfigType("yml")          //文件格式类型
	err = v.ReadInConfig()
	if err != nil {
		log.Fatal("读取配置文件失败：", err)
	}
	return v
}

func dynamicDAGConfig() {
	GlobalDAGConfig.WatchConfig()
	GlobalDAGConfig.OnConfigChange(func(event fsnotify.Event) {
		log.Printf("DAG config change: %s \n", event.String())
	})
	//fmt.Println("Current redis host is: ", dynamic_config.GlobalConfig.GetString("service.redis.host"))
}
