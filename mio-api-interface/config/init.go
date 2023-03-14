package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Gin  Gin  `yaml:"gin"`
	GRPC GRPC `yaml:"grpc"`
}

var Conf Config

func init() {
	yamlFile, err := os.ReadFile("settings.yaml")
	if err != nil {
		fmt.Println("[config init error] os.ReadFile 配置文件读取失败 " + err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		fmt.Println("[config init error] yaml.Unmarshal 配置文件解析失败 " + err.Error())
		return
	}
	fmt.Println("[Success] 配置文件读取成功！！！")
}
