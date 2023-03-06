package tree

import (
	"sort"
)

// Tree 统一定义菜单树的数据结构，也可以自定义添加其他字段
type Tree struct {
	Data     interface{} `json:"data"`     //自定义对象
	Selected bool        `json:"checked"`  //选中
	Children []Tree      `json:"children"` //子节点
}

// ConvertToINodeArray 其他的结构体想要生成菜单树，直接实现这个接口
type INode interface {
	// GetId获取id
	GetID() string
	// 时间戳
	GetTime() int64
	// GetFatherId 获取父id
	GetFatherID() string
	// GetData 获取附加数据
	GetData() interface{}
	// IsRoot 判断当前节点是否是顶层根节点
	IsRoot() bool
}

/*
INodes 节点
*/
type INodes []INode

func (nodes INodes) Len() int {
	return len(nodes)
}

/**
 * Swap
 * 节点排序使用
 */
func (nodes INodes) Swap(i, j int) {
	nodes[i], nodes[j] = nodes[j], nodes[i]
}

/**
 * Less
 * 节点排序使用
 */
func (nodes INodes) Less(i, j int) bool {
	return nodes[i].GetTime() < nodes[j].GetTime()
}

/**
 * GenerateTree
 * 自定义的结构体实现 INode 接口后调用此方法生成树结构
 * nodes 需要生成树的节点
 * selectedNode 生成树后选中的节点
 * menuTrees 生成成功后的树结构对象
 */
func GenerateTree(nodes []INode) (trees []Tree) {
	trees = []Tree{}
	// 定义顶层根和子节点
	var roots, childs []INode

	mapInt := make(map[int]bool)

	for _, v := range nodes {
		if v.IsRoot() {
			// 判断顶层根节点
			roots = append(roots, v)
		} else {
			childs = append(childs, v)
		}
	}

	for index := range childs {
		mapInt[index] = false
	}

	for _, v := range roots {
		childTree := &Tree{
			Data: v.GetData(),
		}
		// 递归
		recursiveTree(childTree, &childs, &mapInt)

		trees = append(trees, *childTree)
	}
	return
}

/**
 * recursiveTree
 * 递归生成树结构
 * tree 递归的树对象
 * nodes 递归的节点
 * selectedNodes 选中的节点
 */
func recursiveTree(tree *Tree, nodes *[]INode, mapInt *map[int]bool) {

	data := tree.Data.(INode)

	for index, v := range *nodes {

		if data.GetID() == v.GetFatherID() && !(*mapInt)[index] {
			childTree := &Tree{
				Data: v.GetData(),
			}
			tree.Children = append(tree.Children, *childTree)
			(*mapInt)[index] = true
			recursiveTree(childTree, nodes, mapInt)

		}

	}
}

/**
 * FindRelationNode
 * 在 allTree 中查询 nodes 中节点的所有父节点
 * nodes 要查询父节点的子节点数组
 * allTree 所有节点数组
 */
func FindRelationNode(nodes, allNodes []INode) (respNodes []INode) {
	nodeMap := make(map[string]INode)
	for _, v := range nodes {
		recursiveFindRelationNode(nodeMap, allNodes, v, 0)
	}

	for _, v := range nodeMap {
		respNodes = append(respNodes, v)
	}
	sort.Sort(INodes(respNodes))
	return
}

/**
 * recursiveFindRelationNode
 * 递归查询关联父子节点
 * nodeMap 查询结果搜集到map中
 * allNodes 所有节点
 * node 递归节点
 * t 递归查找类型：0 查找父、子节点；1 只查找父节点；2 只查找子节点
 */
func recursiveFindRelationNode(nodeMap map[string]INode, allNodes []INode, node INode, t int) {
	nodeMap[node.GetID()] = node
	for _, v := range allNodes {
		if _, ok := nodeMap[v.GetID()]; ok {
			continue
		}
		// 查找父节点
		if t == 0 || t == 1 {
			if node.GetFatherID() == v.GetID() {
				nodeMap[v.GetID()] = v
				if v.IsRoot() {
					// 是顶层根节点时，不再进行递归
					continue
				}
				recursiveFindRelationNode(nodeMap, allNodes, v, 1)
			}
		}
		// 查找子节点
		if t == 0 || t == 2 {
			if node.GetID() == v.GetFatherID() {
				nodeMap[v.GetID()] = v
				recursiveFindRelationNode(nodeMap, allNodes, v, 2)
			}
		}
	}
}
