/*
# @Time : 2020/9/24 21:04
# @Author : smallForest
# @SoftWare : GoLand
*/
package application

import (
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

// 服务器当前时间戳11位
func CurrentTimestamp() int {
	int64num := time.Now().Unix()
	return *(*int)(unsafe.Pointer(&int64num))
}

//随机验证码
func CreateCaptcha() string {
	return fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
}

// 获取随机字符串
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
func Chu(num int) float32 {
	return float32(num) / 100
}

//随机范围正整数
func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max-min) + min
	return randNum
}
