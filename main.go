package main

import "oj/boot"

func main() {
	boot.ViperSetup()
	boot.LoggerSetup()
	boot.MysqlSetup()
	boot.RedisSetup()
	boot.ServerSetup()
}
