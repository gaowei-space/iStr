package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/atotto/clipboard"

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
	var itemStr bool
	var logLevel string

	toTypeMap := map[int]string{
		1: "多行转一行",
		2: "一行转多行",
	}

	flag.IntVar(&toType, "t", 1, "./iStr -t 1|2")
	flag.StringVar(&splitStr, "s", ",", "./iStr -s ,")
	flag.BoolVar(&itemStr, "i", false, "./iStr -i")
	flag.StringVar(&logLevel, "log-level", "info", "./iStr -log-level error")
	flag.Parse()

	level, err := log.ParseLevel(logLevel)
	if err != nil {
		log.Fatal(err)
	}
	log.SetLevel(level)

	// 读取剪切板中的内容到字符串
	fromClipbordStr, err = clipboard.ReadAll()
	if err != nil {
		log.Fatalf("读取剪贴板失败")
	}

	if fromClipbordStr == "" {
		log.Fatalln("剪贴板内容为空")
	}

	_, ok := toTypeMap[toType]
	if !ok {
		log.Fatalln("-t 值输入异常，请输入1[多行转一行]或者2[一行转多行]")
	}

	if toType == 1 {
		arr := strings.Split(fromClipbordStr, "\n")
		var parseArr []string
		for i := 0; i < len(arr); i++ {
			if len(arr[i]) == 0 {
				continue
			}
			parseArr = append(parseArr, strings.TrimSpace(arr[i]))
		}

		if itemStr {
			toClipbordStr = strings.Join(parseArr, fmt.Sprintf("\"%s\"", splitStr))
			toClipbordStr = fmt.Sprintf("\"%s\"", toClipbordStr)
		} else {
			toClipbordStr = strings.Join(parseArr, splitStr)
		}
	} else {
		toClipbordStr = strings.ReplaceAll(fromClipbordStr, splitStr, "\n")
	}

	// 复制内容到剪切板
	clipboard.WriteAll(toClipbordStr)
	fmt.Printf("\n`````````GET````````````\n%s\n\n````````````````````````\n", toClipbordStr)
}
