package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Conf struct {
	Mysql struct {
		Url      string `yaml:"url"`
		UserName string `yaml:"userName"`
		Password string `yaml:"password"`
		DbName   string `yaml:"dbname"`
		Port     string `yaml:"post"`
	}
}

var Config = &Conf{}

func InitConf() {
	yamlFile, err := os.ReadFile("application.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, Config)
	if err != nil {
		fmt.Println(err.Error())
	}
}
