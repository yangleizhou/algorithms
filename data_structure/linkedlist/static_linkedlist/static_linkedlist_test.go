package linkedlist

import (
	"fmt"
	"testing"
)

func TestStaticLinkedlist(t *testing.T) {
	linkedList := Init()
	linkedList.Insert("B", 1)
	linkedList.HeadAdd("A")
	linkedList.Append("D")
	linkedList.Insert("C", 3)
	linkedList.Append("G")
	linkedList.Insert("E", 5)
	linkedList.trave()
	fmt.Println()
	linkedList.Remove(2)
	linkedList.Insert("B", 2)
	linkedList.trave()
	fmt.Println()
	linkedList.Remove(2)
	linkedList.Remove(4)
	linkedList.trave()
	fmt.Println()
	linkedList.destroy()
	linkedList.Insert("B", 1)
	linkedList.HeadAdd("A")
	linkedList.Append("D")
	linkedList.Insert("C", 3)
	linkedList.Append("G")
	linkedList.Insert("E", 5)
	linkedList.trave()
}
