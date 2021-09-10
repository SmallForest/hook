/*
# @Time : 2021/9/8 12:13
# @Author : smallForest
# @SoftWare : GoLand
*/
package utils

import (
	"hook/conf"
	"log"
	"os"
	"strings"
)

// 判断路径PATH的文件是否存在
func FileExist(path string) bool {
	r, err := os.Lstat(path)
	log.Println(r)
	return !os.IsNotExist(err)
}

// IsInWhiteList 判断当前仓库名字是否在白名单
func IsInWhiteList(repository string) bool {
	repository_list := conf.Run().Section("app").Key("repository_list").String()
	return strings.Contains(repository_list, repository)
}

// RepositoryName 获取仓库的名字
func RepositoryName(repository string) string {
	arr := strings.Split(repository, "/")
	if len(arr) == 2 {
		return arr[1]
	}
	return ""
}
