/*
# @Time : 2021/9/8 12:13
# @Author : smallForest
# @SoftWare : GoLand
*/
package utils

import (
	"log"
	"os"
)

// 判断路径PATH的文件是否存在
func FileExist(path string) bool {
	r, err := os.Lstat(path)
	log.Println(r)
	return !os.IsNotExist(err)
}
