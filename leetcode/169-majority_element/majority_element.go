package binary

func majorityElement(nums []int) int {
	res,count :=nums[0],1
	for i:= 1;i<len(nums);i++ {
		if res == nums[i] {
			count++
		}else{
			count--
		}
		if count == 0 {
			res = nums[i]
		}
	}
	return res
} 
