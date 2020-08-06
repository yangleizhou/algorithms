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
	fmt.Printf("0b%b = %d\n", rangeBitwiseAnd(20, 25), rangeBitwiseAnd(20, 25))
	fmt.Printf("0b%b = %d\n", rangeBitwiseAnd(5, 7), rangeBitwiseAnd(5, 7))
	fmt.Printf("0b%b = %d\n", rangeBitwiseAnd(0, 1), rangeBitwiseAnd(0, 1))
}
