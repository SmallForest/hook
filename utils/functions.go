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
	//将仓库拼接成字符串
	repositoryList := strings.Join(conf.Config.Application.Repository, ",")
	return strings.Contains(repositoryList, repository)
}

// RepositoryName 获取仓库的名字
func RepositoryName(repository string) string {
	arr := strings.Split(repository, "/")
	if len(arr) == 2 {
		return arr[1]
	}
	return ""
}

// 判断是否是git命令
func isGitCommand(s string) bool {
	return strings.Contains(s, "git")
}

// GetCommand 获取要执行的命令
func GetCommand(repository string) string {
	name := RepositoryName(repository)
	command := conf.Config.Application.Commands

	//校验需要是git命令才返回
	if isGitCommand(command[name]) {
		return string(command[name])
	}
	return ""
}
