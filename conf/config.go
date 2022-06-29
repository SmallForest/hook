/*
# @Time : 2020/9/21 21:16
# @Author : smallForest
# @SoftWare : GoLand
*/
package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"sync"
)

var once sync.Once

var Config Conf
var err error

type Conf struct {
	Application Application `yaml:"application"`
}

type Application struct {
	Port       string            `yaml:"port"`
	CodeDir    string            `yaml:"code_dir"`
	Repository []string          `yaml:"repository"`
	Commands   map[string]string `yaml:"commands"`
}

//单例
func Run() {
	once.Do(func() {
		log.Println("读取配置文件")
		yamlFile, err := ioutil.ReadFile("/Users/smallforest/GolandProjects/hook/application.yaml")
		if err != nil {
			log.Println(err.Error())
		} // 将读取的yaml文件解析为响应的 struct
		err = yaml.Unmarshal(yamlFile, &Config)
		if err != nil {
			log.Println(err.Error())
		}
	})
}
