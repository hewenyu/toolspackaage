# toolspackage
go 自用的扩展包


go get github.com/hewenyu/toolspackage



## 日志范例

```go
package main

import "github.com/hewenyu/toolspackage/logger"

func main() {

	logger.SetName("test.log")
	logger.SetPath("./logs")

	logger.Info("test 日志")
}

```