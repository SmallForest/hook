/*
# @Time : 2021/9/8 12:13
# @Author : smallForest
# @SoftWare : GoLand
*/
package utils

import (
	"encoding/json"
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

type commandGit string

// 判断是否是git命令
func (g commandGit) isGitCommand() bool {
	return strings.Contains(string(g), "git")
}

// GetCommand 获取要执行的命令
func GetCommand(repository string) string {
	name := RepositoryName(repository)
	command := conf.Run().Section("app").Key("command").String()
	b := []byte(command)
	r := make(map[string]commandGit, 6)
	json.Unmarshal(b, &r)

	//校验需要是git命令才返回
	if r[name].isGitCommand() {
		return string(r[name])
	}
	return ""
}
