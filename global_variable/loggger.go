package global_variable

import "github.com/hewenyu/toolspackage/klog"

const UtilsNames = "log"

var CLOG = klog.NewZap(UtilsNames)
