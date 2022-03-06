# 项目简介
实验性玩具-多行文本转换工具

## 功能特性
* 单行转多行
* 多行转单行

## 支持平台
> Windows、Linux、Mac OS

## 下载
[releases](https://github.com/gaowei-space/iStr/releases)

## 安装

###  二进制安装
1. 解压压缩包  
2. `cd 解压目录`
3. 使用
  * Windows: `iStr.exe`
  * Linux、Mac OS:  `./iStr`

### 源码安装
- 安装Go 1.11+
- `go get github.com/gaowei-space/iStr`
- `export GO111MODULE=on`
- 编译 `make`
- 使用
    * `./bin/iStr`

### 开发
`make build` 编译

`make package` 打包
> 生成当前系统的压缩包 iStr-v1.5-darwin-amd64.tar.gz

`make package-all` 生成Windows、Linux、Mac的压缩包

### 命令
* iStr
    * -i 单个元素是否为字符串，默认 `fasle`
    * -s 指定分隔符，默认 `,`
    * -t 指定转换模式, 1:多行转一行, 2:一行转多行，默认 `1`