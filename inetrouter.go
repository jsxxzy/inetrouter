package inetrouter

import (
	"math/rand"
	"time"
)

// 随机路由最大长度
var maxRouterLength = 18

// GenRandomRouterString 生成路由
//
// 为了安全起见, 每次生成的路由应为一个随机字符串
func GenRandomRouterString() string {
	return createRandString(maxRouterLength)
}

// create by d1y<chenhonzhou@gmail.com>
// 随机函数工具包
// https://github.com/d1y/note/blob/master/utils/randx/randx.go

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// createRandString 创建随机字符(指定长度)
//
// https://stackoverflow.com/a/31832326
func createRandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
