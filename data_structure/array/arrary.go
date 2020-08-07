package array

import (
	"fmt"
	"errors"
)

// Array 数组
type Array struct {
	data   []int
	length uint
}

// NewArray 获取数组对象
func NewArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}
	return &Array{data:make([]int,capacity,capacity)}
}

// Len 获取数组长度
func (this *Array) Len() uint {
	return this.length
}

//isIndexOutOfRange 判断索引是否越界
func (this *Array) isIndexOutOfRange(index uint) bool {
	if index >= uint(cap(this.data)) {
		return true
	}
	return false
}

//Find 通过索引查找数组，索引范围[0,n-1]
func (this *Array) Find(index uint) (int, error) {
	if this.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	return this.data[index], nil
}

//Insert 插入数值到索引index上
func (this *Array) Insert(index uint, v int) error {
	if this.Len() == uint(cap(this.data)) {
		return errors.New("full array")
	}
	if index != this.length && this.isIndexOutOfRange(index) {
		return errors.New("out of index range")
	}

	for i := this.length; i > index; i-- {
		this.data[i] = this.data[i-1]
	}

	this.data[index] = v
	this.length++
	return nil
}

// InsertToTail 往后插入
func (this *Array) InsertToTail(v int) error {
	return this.Insert(this.Len(), v)
}

// Delete 删除索引index上的值
func (this *Array) Delete(index uint) (int, error) {
	if this.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	v := this.data[index]
	for i := index; i < this.Len()-1; i++ {
		this.data[i] = this.data[i+1]
	}
	this.length--
	return v, nil
}

// Print 打印数列
func (this *Array) Print() {
	var format string
	for i := uint(0); i < this.Len(); i++ {
		format += fmt.Sprintf("|%+v", this.data[i])
	}
	fmt.Println(format)
}