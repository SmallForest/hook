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
func update(path string) bool {
	// 锁文件
	lock_path := "/.git/index.lock"
	// 最大次数
	maxTimes := 10
	// 声明times
	times := 0
	// 执行的命令
	command := conf.Run().Section("app").Key("command").String()
	// 检查锁文件是否存在，如果存在则等待1s后再次尝试
	// 当锁文件存在的时候不可以调用git命令会报错 并且只会等待10秒，超过十秒不再执行git命令
	for utils.FileExist(path+lock_path) && times < maxTimes {
		// 当文件存在的时候，执行sleep程序
		time.Sleep(1 * time.Second)
		times++
	}
	if times < maxTimes {
		sh.NewSession().SetDir(path).Command("bash", "-c", command).Run()
	} else {
		log.Println("超时未执行，此时times ：", times)
	}
	return true
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

		code_dir := conf.Run().Section("app").Key("code_dir").String()
		log.Println("代码运行的路径是", code_dir)

		// 注意 Unable to create '/usr/share/nginx/html/code/jinlianlian-app-api/.git/index.lock': File exists.
		// 切换目录 ....
		// 识别到的仓库名字在白名单里面，进行下一步操作
		if utils.IsInWhiteList(repository) {
			log.Println("仓库", repository, "在白名单中。执行更新")
			if name := utils.RepositoryName(repository); name != "" {
				update(code_dir + name)
				log.Println("更新成功")
			}else{
				log.Println("更新失败")
			}
		} else {
			log.Println("仓库", repository, "不在白名单中。不执行更新")
		}

	})
	_ = router.Run(conf.Run().Section("app").Key("start_listen_port").String())
}
