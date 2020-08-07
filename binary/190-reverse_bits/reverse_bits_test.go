package binary

import (
	"fmt"
	"testing"
)

func TestReverseBits(t *testing.T) {
	fmt.Printf("%032b\n%032b\n", 0B00000010100101000001111010011100, reverseBits(0B00000010100101000001111010011100))

	fmt.Println()

	fmt.Printf("%032b\n%032b\n", 0B11111111111111111111111111111101, reverseBits(0B11111111111111111111111111111101))

}
