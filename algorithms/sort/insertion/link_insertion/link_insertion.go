package insertion

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
	for i := 1; i <= len(nums); i++ {
		node = Node{rc: nums[i-1], next: 0}
		nl = append(nl, node)
	}
	return nl
}

func (nl NodeList) sort() {
	var low, high int
	for i := 2; i < len(nl); i++ {
		low, high = 0, nl[0].next
		for nl[high].rc < nl[i].rc {
			low = high
			high = nl[high].next
		}
		nl[low].next = i
		nl[i].next = high
	}
}

func (nl NodeList) arrange() {
	p := nl[0].next //p指示第一个记录的当前位置，即还没有排序的最小值
	for i := 1; i < len(nl); i++ {
		for p < i { //i之前的都是排好的
			p = nl[p].next
		}
		q := nl[p].next //下一个要排序的记录，p排好就到q了，找第i+1个记录做准备
		if p != i {     //交换记录，使第i个记录到位
			nl[p].rc, nl[i].rc = nl[i].rc, nl[p].rc
			nl[p].next = nl[i].next
			nl[i].next = p //指向被移走的记录（搬家的记录）
		}
		p = q //找第i+1个记录做准备
	}
}

func linkInsertion(nums []int) []int {
	nl := Init(nums)
	nl.sort()
	nl.arrange()
	for i := 1; i < len(nl); i++ {
		nums[i-1] = nl[i].rc
	}
	return nums
}
