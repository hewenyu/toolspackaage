package network

import "net"

// GetLocalAddress 获取本地的ip地址
func GetLocalAddress() []string {
	var _add = make([]string, 0)
	
	adders, err := net.InterfaceAddrs()
	if err != nil {
		return _add
	}
	
	for _, addr := range adders {
		
		if inet, ok := addr.(*net.IPNet); ok && !inet.IP.IsLoopback() {
			if inet.IP.To4() != nil {
				
				_add = append(_add, inet.IP.String())
			}
		}
	}
	return append(_add[:0:0], _add...)
}
