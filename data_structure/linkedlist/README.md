## 链表 linkedlist

- [定义](#1)
- [链表的结构](#2)
- [链表优缺点](#3)



- ## <i id="1"></i>**`定义`**  
`链表`是一种线性的数据结构,和[数组](../array)不同的是,`链表`不需要一块联系的内存空间,它通过指针将一组零散的内存空间串联起来,是`链式存储`结构一种。

- ## <i id="2"></i>**`链表结构`**  
    `链表`的结构非常多,常见的有如下：</br>  
    
    - [单链表](./single_linkedlist)
    - [双向链表](./double_linkedlist)
    - [循环链表](./circular_linkedlist)
        - [单向循环链表](./circular_linkedlist/single_circular_linkedlist)
        - [双向循环链表](./circular_linkedlist/double_circular_linkedlist)
    - [静态链表](./static_linkedlist)
    - [十字链表](./orthogonal_list)
    - [跳跃表](./skip_list)


- ## <i id="3"></i>**`优缺点`**  
    - `优点`  
        1. `链表`能灵活分配内存空间。</br>
        2. 能在O(1)时间内删除或者添加元素。删除情况下O(1)前提该元素的前一个元素已知。 </br>

    - `缺点`  
        1. 不能像[数组](../array)能通过下标迅速读取元素,每次都要从链表头开始一个一个读取。</br> 
        2. 查询第K个元素需要O(k)时间。</br>

    - `复杂度对比`  
    时间复杂度 | 数组 | 链表   
    ---- | ---- | ----   
    插入删除| O(n) | O(1)  
    随机访问| O(1) | O(n)  

    
     

