package binary

import (
	"testing"
)

func TestMaxProduct(t *testing.T) {
	words := []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}
	maxProduct(words)

	words = []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}
	maxProduct(words)

	words = []string{"a", "aa", "aaa", "aaaa"}
	maxProduct(words)
}
