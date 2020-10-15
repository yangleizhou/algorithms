package stack

import (
	"fmt"
	"testing"
)

func TestStackUseDecimalToAny(t *testing.T) {
	fmt.Println(10)
	DecimalToAny(10, 10)
	fmt.Println()
	fmt.Println()

	fmt.Println(10)
	DecimalToAny(10, 2)
	fmt.Println()
	fmt.Println()

	fmt.Println(10)
	DecimalToAny(10, 8)
	fmt.Println()
	fmt.Println()

	fmt.Println(10)
	DecimalToAny(10, 16)
	fmt.Println()
	fmt.Println()
}
