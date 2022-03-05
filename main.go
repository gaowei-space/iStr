package main

import (
	"flag"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

func main() {
	// 一行转多行
	// 参数：分割符号
	// 多行转一行
	// 参数：分割符号
	// 分割类型：数值，字符串

	var fromClipbordStr string
	var toClipbordStr string
	var toType int
	var splitStr string
	var itemToStr bool
	var logLevel string

	toTypeMap := map[int]string{
		1: "多行转一行",
		2: "一行转多行",
	}

	flag.IntVar(&toType, "t", 1, "./iStr -t 1|2")
	flag.StringVar(&splitStr, "s", ",", "./iStr -s ,")
	flag.BoolVar(&itemToStr, "i", false, "./iStr -i")
	flag.StringVar(&logLevel, "log-level", "info", "-log-level error")
	flag.Parse()

	level, err := log.ParseLevel(logLevel)
	if err != nil {
		log.Fatal(err)
	}
	log.SetLevel(level)

	_, ok := toTypeMap[toType]
	if !ok {
		log.Fatalf("type 值输入异常，请输入1[多行转一行]或者2[一行转多行]")
	}

	if toType == 1 {
		fromClipbordStr = `
	slack
	slack
	mail
	mail
	webhook
	webhook`
		if itemToStr {
			fromClipbordStr = strings.ReplaceAll(fromClipbordStr, " ", "")
			fromClipbordStr = strings.ReplaceAll(fromClipbordStr, "\n", fmt.Sprintf("\"%s\"", splitStr))
			toClipbordStr = fmt.Sprintf("\"%s\"", fromClipbordStr)
		} else {
			toClipbordStr = strings.ReplaceAll(fromClipbordStr, "\n", splitStr)
		}
	} else {
		fromClipbordStr = "slack,hello,ddd"
		toClipbordStr = strings.ReplaceAll(fromClipbordStr, splitStr, "\n")
	}

	fmt.Println(toClipbordStr)
}
