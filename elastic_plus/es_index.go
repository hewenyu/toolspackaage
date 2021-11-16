package elastic_plus

// ESIndex 索引的数据格式
type ESIndex struct {
	UUID   string `json:"uuid"`
	Health string `json:"health"`
	Status string `json:"status"`
	Index  string `json:"index"`
}

// ESIndexList 索引List
type ESIndexList []ESIndex
