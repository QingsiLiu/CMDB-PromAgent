package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"promagent/config"
)

func initConfig(path string) *config.AgentConfig {
	var config config.AgentConfig

	viper.AutomaticEnv()
	viper.SetEnvPrefix("PROM_AGENT")

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	fmt.Println(config)
	return &config
}

func main() {

	var (
		verbose bool
		help    bool
		path    string
	)

	flag.BoolVar(&verbose, "verbose", false, "verbose")
	flag.BoolVar(&help, "help", false, "help")
	flag.StringVar(&path, "path", "./etc/promagent.yaml", "config path")

	flag.Usage = func() {
		fmt.Println("usage: promagent [-verbose] [-config]")
		flag.PrintDefaults()
	}

	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	// 初始化配置
	config := initConfig(path)
	fmt.Println(config)
	// 初始化日志
	// 启动

}
