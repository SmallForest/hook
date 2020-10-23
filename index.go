/*
# @Time : 2019-07-22 09:18
# @Author : smallForest
# @SoftWare : GoLand
*/
package main

import (
	"fmt"
	_ "github.com/astaxie/beego/httplib"
	"github.com/bitly/go-simplejson"
	"github.com/codeskyblue/go-sh"
	"github.com/gin-gonic/gin"
	_ "github.com/gogap/wechat/util"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"hook/conf"
	"io/ioutil"
	_ "time"
)

func main() {
	// 发布模式开启
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// 以上都是测试

	router.POST("/hook", func(c *gin.Context) {
		res := c.Request.Body
		fmt.Println("hook结果", res)
		bodydata, err := ioutil.ReadAll(res)
		if err != nil {
			fmt.Println(err)
		}
		j, err := simplejson.NewJson(bodydata)
		if err != nil {
			fmt.Printf("err", err)
		}
		repository, err := j.Get("repository").Get("full_name").String()
		if err != nil {
			fmt.Printf("err", err)
		}
		fmt.Println("识别到的仓库是", repository)

		clone_url, err := j.Get("repository").Get("clone_url").String()
		if err != nil {
			fmt.Printf("err", err)
		}
		fmt.Println("识别到的clone_url是", clone_url)
		// 切换目录 ....
		if repository == "smallForest/jinlianlian-platform-api" {
			fmt.Println("1平台端API")
			sh.NewSession().SetDir("/usr/share/nginx/html/code/jinlianlian-platform-api").Command("git", "checkout", "develop").Command("git", "pull").Run()
		}
		if repository == "smallForest/jinlianlian-app-api" {
			fmt.Println("2APP端API")
			sh.NewSession().SetDir("/usr/share/nginx/html/code/jinlianlian-app-api").Command("git", "checkout", "develop").Command("git", "pull").Run()
		}
		if repository == "smallForest/jinlianlian-business-api" {
			fmt.Println("3商家端API")
			sh.NewSession().SetDir("/usr/share/nginx/html/code/jinlianlian-business-api").Command("git", "checkout", "develop").Command("git", "pull").Run()
		}
		if repository == "smallForest/jinlianlian-business-doc" {
			fmt.Println("4商家端文档")
			sh.NewSession().SetDir("/usr/share/nginx/html/code/jinlianlian-business-doc").Command("git", "checkout", "develop").Command("git", "pull").Run()
		}
		if repository == "smallForest/jinlianlian-app-doc" {
			fmt.Println("5APP端文档")
			sh.NewSession().SetDir("/usr/share/nginx/html/code/jinlianlian-app-doc").Command("git", "checkout", "develop").Command("git", "pull").Run()
		}
		if repository == "smallForest/jinlianlian-platform-doc" {
			fmt.Println("6平台端文档")
			sh.NewSession().SetDir("/usr/share/nginx/html/code/jinlianlian-platform-doc").Command("git", "checkout", "develop").Command("git", "pull").Run()
		}
		if repository == "smallForest/jinlianlian-business-web" {
			fmt.Println("7商家web后台")
			sh.NewSession().SetDir("/usr/share/nginx/html/code/jinlianlian-business-web").Command("git", "checkout", "develop").Command("git", "pull").Run()
		}

	})
	_ = router.Run(conf.Run().Section("app").Key("start_listen_port").String())
}
