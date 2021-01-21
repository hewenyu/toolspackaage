package main

import "github.com/hewenyu/toolspackage/logger"

func main() {
	logger.SetName("test111.log")
	logger.SetPath("./logs")
	logger.Info("hello")
}
