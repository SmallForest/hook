/*
# @Time : 2019-07-22 09:18
# @Author : smallForest
# @SoftWare : GoLand
*/
package main

import (
	"github.com/bitly/go-simplejson"
	"github.com/codeskyblue/go-sh"
	"github.com/gin-gonic/gin"
	"hook/conf"
	"hook/utils"
	"io/ioutil"
	"log"
	"time"
	_ "time"
)

// 以上都是测试
func update(path string) {
	// 锁文件
	lock_path := "/.git/index.lock"
	// 最大次数
	maxTimes := 10
	// 声明times
	times := 0
	// 检查锁文件是否存在，如果存在则等待1s后再次尝试
	// 当锁文件存在的时候不可以调用git命令会报错 并且只会等待10秒，超过十秒不再执行git命令
	for utils.FileExist(path+lock_path) && times < maxTimes {
		// 当文件存在的时候，执行sleep程序
		time.Sleep(1 * time.Second)
		times++
	}
	if times < maxTimes {
		sh.NewSession().SetDir(path).Command("bash", "-c", "git checkout develop && git pull").Run()
	} else {
		log.Println("超时未执行，此时times ：", times)
	}
}

func main() {
	// 发布模式开启
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/hook", func(c *gin.Context) {
		res := c.Request.Body
		log.Println("hook结果", res)
		bodydata, err := ioutil.ReadAll(res)
		if err != nil {
			log.Println(err)
		}
		j, err := simplejson.NewJson(bodydata)
		if err != nil {
			log.Printf("err %v", err)
		}
		repository, err := j.Get("repository").Get("full_name").String()
		if err != nil {
			log.Printf("err %v", err)
		}
		log.Println("识别到的仓库是", repository)

		clone_url, err := j.Get("repository").Get("clone_url").String()
		if err != nil {
			log.Printf("err %v", err)
		}
		log.Println("识别到的clone_url是", clone_url)

		// 注意 Unable to create '/usr/share/nginx/html/code/jinlianlian-app-api/.git/index.lock': File exists.
		// 切换目录 ....
		if repository == "smallForest/jinlianlian-platform-api" {
			log.Println("1平台端API")
			update("/usr/share/nginx/html/code/jinlianlian-platform-api")
		} else if repository == "smallForest/jinlianlian-app-api" {
			log.Println("2APP端API")
			update("/usr/share/nginx/html/code/jinlianlian-app-api")
		} else if repository == "smallForest/jinlianlian-business-api" {
			log.Println("3商家端API")
			update("/usr/share/nginx/html/code/jinlianlian-business-api")
		} else if repository == "smallForest/jinlianlian-business-doc" {
			log.Println("4商家端文档")
			update("/usr/share/nginx/html/code/jinlianlian-business-doc")
		} else if repository == "smallForest/jinlianlian-app-doc" {
			log.Println("5APP端文档")
			update("/usr/share/nginx/html/code/jinlianlian-app-doc")
		} else if repository == "smallForest/jinlianlian-platform-doc" {
			log.Println("6平台端文档")
			update("/usr/share/nginx/html/code/jinlianlian-platform-doc")
		} else if repository == "smallForest/jinlianlian-business-web" {
			log.Println("7商家web后台")
			update("/usr/share/nginx/html/code/jinlianlian-business-web")
		} else if repository == "smallForest/jinlianlian-platform-web" {
			log.Println("8平台苍穹web后台")
			update("/usr/share/nginx/html/code/jinlianlian-platform-web")
		} else if repository == "smallForest/jinlianlian-app-h5" {
			log.Println("9H5web")
			update("/usr/share/nginx/html/code/jinlianlian-app-h5")
		}

	})
	_ = router.Run(conf.Run().Section("app").Key("start_listen_port").String())
}
