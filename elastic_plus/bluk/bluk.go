package bluk

// InsertBody 插入es 的数据
type InsertBody struct {
	ID   string     `json:"id"`
	Body InsertJson `json:"body"`
}

type BulkData struct {
	Index string // 插入的索引
	Data  InsertBody
}

type BulkDataList []BulkData
