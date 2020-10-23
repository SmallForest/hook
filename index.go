/*
# @Time : 2019-07-22 09:18
# @Author : smallForest
# @SoftWare : GoLand
*/
package main

import (
	"fmt"
	_ "github.com/astaxie/beego/httplib"
	"github.com/codeskyblue/go-sh"
	"github.com/gin-gonic/gin"
	_ "github.com/gogap/wechat/util"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"hook/conf"
	"io/ioutil"
	_ "time"
)

func main() {

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// 以上都是测试

	router.POST("/hook", func(c *gin.Context) {
		res := c.Request.Body
		fmt.Println("回调结果", res)
		bodydata, err := ioutil.ReadAll(res)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(bodydata))
		//执行git 命令
		sh.Command("git", "add ./").Run()
	})

	_ = router.Run(conf.Run().Section("app").Key("start_listen_port").String())
}
