package insertion

// http://c.biancheng.net/view/3442.html
// https://blog.csdn.net/zhangkaihang/article/details/7394806
// https://blog.csdn.net/qq_36183935/article/details/80217487
// https://www.cnblogs.com/ciyeer/p/9075303.html

// IntMax 最大值
const IntMax = int(^uint(0) >> 1)

// Node 结点
type Node struct {
	rc   int //记录项
	next int //指针域，由于在数组中，所有需要记录下一个结点所在数组位置的下标即可
}

// NodeList  链表
type NodeList []Node

func Init(nums []int) NodeList {
	var nl NodeList
	node := Node{IntMax, 1}
	nl = append(nl, node)
	for i := 1; i < len(nums); i++ {
		node = Node{nums[i-1], 0}
		nl = append(nl, node)
	}
	return nl
}

func (nl NodeList) sort() {

}

func linkInsertion(nums []int) []int {
	nl := Init(nums)
}
