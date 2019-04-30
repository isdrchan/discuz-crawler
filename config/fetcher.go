package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type CrawlerConfig struct {
	Selector struct {
		Forum   string `yaml:"forum"`
		List    string `yaml:"list"`
		Content string `yaml:"content"`
	}
}

var Crawler = CrawlerConfig{}

func init() {
	yamlFile, err := ioutil.ReadFile("c:\\config.yaml")
	if err != nil {
		log.Fatalf("读取yaml配置文件失败: %s ", err)
	}
	err = yaml.Unmarshal(yamlFile, &Crawler)
	if err != nil {
		log.Fatalf("yaml配置文件格式有误: %s", err)
	}
}
