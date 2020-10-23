/*
# @Time : 2020/9/21 21:16
# @Author : smallForest
# @SoftWare : GoLand
*/
package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"sync"
)

var once sync.Once

var cfg *ini.File
var err error

//单例
func Run() *ini.File {
	once.Do(func() {
		fmt.Println("读取配置文件")
		cfg, err = ini.Load("application.ini")
		if err != nil {
			fmt.Printf("Fail to read file: %v", err)
			os.Exit(1)
		}
	})

	return cfg
	// fmt.Println("App Mode:", cfg.Section("").Key("app_mode").String())
	// fmt.Println("Data Path:", cfg.Section("paths").Key("data").String())
	// fmt.Println("port:", cfg.Section("server").Key("http_port").MustInt(3306))
}
