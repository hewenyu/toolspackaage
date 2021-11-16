package bluk


// InsertJson 插入数据格式化
type InsertJson interface {
	GetJson() ([]byte, error)
}

