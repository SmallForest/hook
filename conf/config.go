/*
# @Time : 2020/9/21 21:16
# @Author : smallForest
# @SoftWare : GoLand
*/
package conf

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
	"sync"
)

var once sync.Once

var cfg *ini.File
var err error

//单例
func Run() *ini.File {
	once.Do(func() {
		log.Println("读取配置文件")
		cfg, err = ini.Load("application.ini")
		if err != nil {
			log.Printf("Fail to read file: %v", err)
			os.Exit(1)
		}
	})

	return cfg
	// log.Println("App Mode:", cfg.Section("").Key("app_mode").String())
	// log.Println("Data Path:", cfg.Section("paths").Key("data").String())
	// log.Println("port:", cfg.Section("server").Key("http_port").MustInt(3306))
}
