package tree

import "time"

/*
DataNode 查询json
*/
type DataNode struct {
	Name     string    `json:"name"`
	Mydate   time.Time `json:"mydate"`
	BaseID   string    `json:"id"`
	FatherID string    `json:"father"`
}

/*
DataNodeList 返回给前端的数据信息整合列表
*/
type DataNodeList []DataNode

/*
GetID 获取ID
*/
func (s DataNode) GetID() string {
	return s.BaseID
}

/*
GetFatherID 父id
*/
func (s DataNode) GetFatherID() string {
	return s.FatherID
}

/*
GetTime 时间
*/
func (s DataNode) GetTime() int64 {
	return s.Mydate.UnixNano()
}

/**
 * GetData
 * 获取数据
 */
func (s DataNode) GetData() interface{} {
	return s
}

/**
 * IsRoot
 * 判断是否是根目录
 */
func (s DataNode) IsRoot() bool {
	// 这里通过FatherId等于0 或者 FatherId等于自身Id表示顶层根节点
	return s.FatherID == "0" || s.FatherID == s.BaseID
}

// ConvertToINodeArray 将当前数组转换成父类 INode 接口 数组
func (s DataNodeList) ConvertToINodeArray() (nodes []INode) {
	for _, v := range s {
		nodes = append(nodes, v)
	}
	return
}

/**
 * SortByID
 * 按照结果排序
 */
func (c *DataNodeList) SortByID() (final []Tree) {

	final = GenerateTree(DataNodeList.ConvertToINodeArray(*c))

	return
}
