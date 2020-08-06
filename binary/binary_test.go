package binary

import (
	"fmt"
	"testing"
)

func TestParity(t *testing.T) {
	parity(10)
	parity(1)
	parity(0)
	parity(3)
}

func TestConvertLow(t *testing.T) {
	var str = []byte{'A', 'B', 'C', 'E', 'F', 'G', 'd', 'g', 'h'}
	fmt.Println("source =  ", string(str))
	convertLower(str)
	fmt.Println("after =  ", string(str))
}

func TestConvertSupper(t *testing.T) {
	var str = []byte{'A', 'B', 'C', 'E', 'F', 'G', 'd', 'g', 'h'}
	fmt.Println("source =  ", string(str))
	convertSupper(str)
	fmt.Println("after =  ", string(str))
}

func TestConvertLowerSupper(t *testing.T) {
	var str = []byte{'A', 'B', 'C', 'E', 'F', 'G', 'd', 'g', 'h'}
	fmt.Println("source =  ", string(str))
	convertLowerSupper(str)
	fmt.Println("after =  ", string(str))
}

func TestIsSameSymbol(t *testing.T) {
	var x, y, z = 1, 2, -1
	isSameSymbol(x, y, z)
}

func TestSwap(t *testing.T) {
	swap(1, 2)
}

func TestRemoveLast1(t *testing.T) {
	if removeLast1(0B10011) != 0B10010 {
		t.Errorf("result should is 0B10010")
	}
	if removeLast1(0B10010) != 0B10000 {
		t.Errorf("result should is 0B10000")
	}
}

func TestGetLast1(t *testing.T) {
	if getLast1(0B10010) != 0B10 {
		t.Errorf("result should is 0B10")
	}
	if getLast1(0B10011) != 0B1 {
		t.Errorf("result should is 0B1")
	}
}

func TestIsPowerOf2(t *testing.T) {
	if !isPowerOf2(5) {
		fmt.Printf("%d 不是2的指数幂\n", 5)
	}
	if isPowerOf2(8) {
		fmt.Printf("%d 是2的指数幂\n", 8)
	}
}

func TestSingleNumber(t *testing.T) {
	fmt.Println(singleNumber(15, 6, 6, 9, 9, 10, 10, 20, 20))
}

func TestSingleNumber2(t *testing.T) {
	fmt.Println(singleNumber2(2, []int{15, 6, 6, 9, 9, 10, 10, 20, 20}))
	fmt.Println(singleNumber2(3, []int{15, 6, 6, 6, 9, 9, 9, 10, 10, 10, 20, 20, 20}))
	fmt.Println(singleNumber2(4, []int{15, 6, 6, 6, 6, 10, 10, 10, 10, 20, 20, 20, 20}))
}

func TestHammingWeight(t *testing.T) {
	hammingWeight(0B0, 0B10111)
	hammingWeight(0B1001, 0B10111)
}
