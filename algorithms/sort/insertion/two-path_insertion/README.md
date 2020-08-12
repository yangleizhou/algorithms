## 2-路插入排序 2-path insertion

- **`定义`**  
`2-路插入排序算法`是在[折半插入排序](../binary_insertion)的基础上对其进行改进,减少其在排序过程中移动记录的次数从而提高效率。</br>


- **`过程简单描述`**  
1. 构建一相同大小的`循环数组temp`,把原数组的元素依次。</br>
    1. 把原数组第一个值nums[0]赋值给temp[0],即temp[0] = nums[0],作为循环数组的第一个数,当然也可以选择其它的数作为第一个数。定义变量`first,final,n:=0,0,len(nums)`,`first,final` 作为temp数组最小/最大值索引。</br>
    2. 待插入元素比最小元素小,`nums[i] `<` temp[first]`,变化`first`和`temp[first]`, `first = (first -1 + n)%n`,`temp[first] = nums[i] `。</br>
    3. 待插入元素大于等于最大元素大, `nums[i] `>=` temp[final]`,变化`final`和`temp[last]`,`last = (last+1+n)%n`,`temp[last]=nums[i]`。</br>
    4. 待插入元素大于等于最小元素,小于最大元素，即 temp[first]`<=` nums[i] && nums[i] `<` temp[last],则temp在[first,last]区间元素都是有序的,选择适当的策略插入。</br>
2. 最后按合适次序赋值回原数组。</br>