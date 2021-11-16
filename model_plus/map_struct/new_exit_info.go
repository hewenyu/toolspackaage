package map_struct

// NewExitInfo 初始化
func NewExitInfo() *ExitInfo {
	return &ExitInfo{
		data: make(map[string]struct{}),
	}
}
