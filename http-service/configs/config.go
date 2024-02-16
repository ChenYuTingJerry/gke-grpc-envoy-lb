package configs

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

const (
	fileType = "yaml"
)

var (
	C        = new(Config)
	filePath = flag.String("file", "", "config file path (yaml)")
)

type Config struct {
	App       AppConfig
	Component ComponentConfig
}

func InitConfigs() {
	fmt.Println("Config init")
	flag.Parse()
	if *filePath != "" {
		fromFile()
	}

	fmt.Println("C.App = ", C.App)
	fmt.Println("C.Comp = ", C.Component)
}

func registerFlags() {
	appFlags()
	componentFlags()
}

func fromFile() {
	viper.SetConfigType(fileType)
	viper.SetConfigFile(*filePath)
	err := viper.ReadInConfig()
	if err != nil {
		panic("Read configs error，reason：" + err.Error())
	}

	err = viper.Unmarshal(&C)
	if err != nil {
		panic("Unmarshal error: " + err.Error())
	}
}
