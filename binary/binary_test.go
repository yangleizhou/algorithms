package binary

import (
	"fmt"
	"testing"
)

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

func TestIsPowerOf2(t *testing.T) {
	if !isPowerOf2(5) {
		fmt.Printf("%d 不是2的指数幂\n", 5)
	}
	if isPowerOf2(8) {
		fmt.Printf("%d 是2的指数幂\n", 8)
	}
}

func TestHammingWeight(t *testing.T) {
	hammingWeight(0B1001, 0B10111)
	hammingWeight(0B0, 0B10111)
}
