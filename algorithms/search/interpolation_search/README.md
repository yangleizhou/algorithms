## 插值查找算法 interpolation search  

- **`定义`**  
`插值查找`，`有序表`的一种查找方式。`插值查找`是根据查找关键字与查找表中最大最小记录关键字比较后的查找方法。插值查找基于[二分查找](../binary_search)优化，将查找点的选择改进为自适应选择，提高查找效率。</br>
获取到的关键字对于索引为最左索引  
```go
var nums = []int{1, 2, 2, 2, 2, 2, 2, 2, 5, 5, 5, 5, 7, 9, 23, 45, 67} 
var target int = 2
index :=interpolationSearch(nums,target)  
获取到index 为1
```


- **`算法来源`**  
插值查找,将[二分查找](../binary_search) mid = low + ((high-low)>>1) 改为 mid = low + (high - low)*(target - a[low])/(a[high] - a[low]),根据查找关键字target 与查找表中最大最小记录关键字比较后的查找方法,核心在于插值计算公式:  
(target - a[low]) /  (a[high] - a[low])  (target 待查找关键字,数组a有序表)  



- **`过程简单描述`**
1. 判断target,是否在[a[low],a[high]]区间内,是的话根据公式获取mid,否则返回-1。 </br>
1. 在有序表中，取val = a[mid]作为比较关键字，若给定值与val相等，则查找成功。</br>
2. 若给定的值小于val，则在mid左半区间继续查找；若给定值大于val，则在mid右半区间继续查找。</br>
3. 不断重复这个过程，直到查找成功。否则查找失败。</br>  


- **`复杂度`**      
O(logn)  
[二分查找时间复杂度](../binary_search#binary_search_space) 时间复杂度O(logn),插值查询跟[二分查找](../binary_search)不同的地方是 1/2系数。

- **`适用`**  
对于数据量较大，关键字`分布比较均匀`的`有序表`来说，采用插值查找, 速度较快.



