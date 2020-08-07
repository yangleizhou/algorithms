# [169. Majority Element](https://leetcode.com/problems/majority-element/)


## 题目:

Given an array of size n, find the majority element. The majority element is the element that appears **more than** `⌊ n/2 ⌋` times.

You may assume that the array is non-empty and the majority element always exist in the array.

**Example 1:**

    Input: [3,2,3]
    Output: 3

**Example 2:**

    Input: [2,2,1,1,1,2,2]
    Output: 2

## 题目大意


给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。你可以假设数组是非空的，并且给定的数组总是存在众数。


## 解题思路

- 题目要求找出数组中出现次数大于 `⌊ n/2 ⌋` 次数。要求空间复杂度为 O(1)。简单题。  
    - 利用摩尔投票 维护一个序列元素和一个计数器 count 
        - 序列元素跟下一个元素相等，count++,否则count--,当count == 0 ,重新赋值序列元素;由于数组中出现次数大于 `⌊ n/2 ⌋` 次数,最终count>0,对于的序列元素就是需要寻找的
