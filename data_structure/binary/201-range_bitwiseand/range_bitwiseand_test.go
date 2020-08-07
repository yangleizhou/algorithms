package binary

import (
	"fmt"
	"testing"
)

func TestRangeBitwiseAnd(t *testing.T) {
	fmt.Printf("0b%b = %d\n", rangeBitwiseAnd1(1000, 1500), rangeBitwiseAnd1(1000, 1500))
	fmt.Printf("0b%b = %d\n", rangeBitwiseAnd1(20, 25), rangeBitwiseAnd1(20, 25))
	fmt.Printf("0b%b = %d\n", rangeBitwiseAnd1(5, 7), rangeBitwiseAnd1(5, 7))
	fmt.Printf("0b%b = %d\n", rangeBitwiseAnd1(0, 1), rangeBitwiseAnd1(0, 1))
	fmt.Println()
	fmt.Printf("0b%b = %d\n", rangeBitwiseAnd2(20, 25), rangeBitwiseAnd2(20, 25))
	fmt.Printf("0b%b = %d\n", rangeBitwiseAnd2(5, 7), rangeBitwiseAnd2(5, 7))
	fmt.Printf("0b%b = %d\n", rangeBitwiseAnd2(0, 1), rangeBitwiseAnd2(0, 1))
}

func TestRangeBitwiseAndCmp1(t *testing.T) {
	rangeBitwiseAnd1(1000, 262143)
}

func TestRangeBitwiseAndCmp2(t *testing.T) {
	rangeBitwiseAnd2(1000, 262143)
}
