package array

import "testing"

func TestArray(t *testing.T) {
	arr := NewArray(10)
	arr.InsertToTail(20)
	arr.InsertToTail(20)
	arr.Insert(1, 30)
	arr.Print()

	arr.Delete(0)
	arr.Print()
}
