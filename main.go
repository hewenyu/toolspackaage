package toolspackage

import "github.com/hewenyu/toolspackage/logger"

func main() {
	var a *logger.Hliog
	a.SetName("test")
	a.SetPath("./logs")
}
