# toolspackage
go 自用的扩展包


go get github.com/hewenyu/toolspackage@v0.0.8



## 日志范例

```go
package main

import "github.com/hewenyu/toolspackage/klog"


func main() {
	klog.Zap().Info("test 日志")
}

```