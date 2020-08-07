package binary

import "fmt"

func maxProduct(words []string) int {
	length := len(words)
	maxProduct := 0
	values := make([]int, length)
	for i := 0; i < length; i++ {
		values[i] = 0
		for index := range words[i] {
			values[i] |= (1 << (words[i][index] - 'a'))
		}
	}
	var index1, index2 = -1, -1
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if (values[i]&values[j]) == 0 && len(words[i])*len(words[j]) > maxProduct {
				maxProduct = len(words[i]) * len(words[j])
				index1 = i
				index2 = j
			}
		}
	}
	if index1 != -1 && index2 != -1 {
		fmt.Printf("len(%s) * len(%s) = %d\n", words[index1], words[index2], maxProduct)
	}
	return maxProduct
}
