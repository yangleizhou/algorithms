## 静态链表 static list

- [定义](#1)
- [实现原理](#2)
- [过程简单描述](#3)
- [优点](#4)
- [缺点](#5)


- ## <i id="1"></i>**`定义`**  
`静态链表`,用[数组](../../array)来实现链式存储结构。兼顾`顺序表`和链表的优点，是`顺序表`和`链表`的升级,`静态链表`的数据全部存储在数组中(顺序表)，但存储的位置是随机的，数据直接的一对一关系是通过`游标`维持。</br>
 结点：数据域,用于存储数据元素的值。  
>>  游标, 即数组下标，表示直接后继元素所在数组中的位置。  

- ## <i id="2"></i>**`实现原理`**  
1. 使用结构体数组,结构体有`游标cursor` 和`数据域 data`
2. 一个数组分量表示一个结点，用cursor代替指针指示结点在数组中的相对位置 
3. 数组逻辑上分为两个链表：`备用链表`(空闲的结点：未被使用或已被删除的分量)和`数据链表`(已被使用的结点)
4. 对于数组的第一个和最后一个元素做特殊处理，它们的data不存放数据
5. 数组的第一个元素，即下标为0的cursor 存放`备用链表`的第一个结点下标
6. 数组的最后一个元素，即下标MAXSIZE-1的cursor,存放第一个有数值的元素的下标，相当于[单链表](../single_linkedlist)中的`首元结点`，初始化时没有数据域，此时指向0
7. 最后一个数据的游标cursor指向0,即最后的数据链表游标对应备用链表第一个结点下标的游标值  

- ## <i id="3"></i>**`过程简单描述`**   
1. 静态链表初始化，根据[实现原理](#2)的第3、4、5、6点,创建`备用链表` 

    | |   |   |   |   |   |   |     |   |
    ---- |----| ----|----| ---- | ---- |----|-----|-----|
    游标 | 1 | 2 | 3 | 4 | 5 | 6 | ... | 0 |
    数据 |   |   |   |   |   |   |     |   |
    下标 | 0 | 1 | 2 | 3 | 4 | 5 | ... | 999 |  

```go
func Init(L Linkedlist) []Node {}
```

2. 插入,每当进行插入时，从`备用链表`取得第一个结点作为待插入的新结点,在A之后插入B。
   
    | |   |   |   |   |   |   |     |   |
    ---- |----| ----|----| ---- | ---- |----|-----|-----|
    游标 | 5 | 2 | 3 | 4 | 0 | 6 | ... | 1 |
    数据 |   | A | C | D | E |   |     |   |
    下标 | 0 | 1 | 2 | 3 | 4 | 5 | ... | 999 | 
 
    </br></br>

    
    | |   |   |   |   |   |   |     |   |
    ---- |----| ----|----| ---- | ---- |----|-----|-----|
    游标 | 6 | 5 | 3 | 4 | 0 | 2 | ... | 1 |
    数据 |   | A | C | D | E | B |     |   |
    下标 | 0 | 1 | 2 | 3 | 4 | 5 | ... | 999 | 

3. 删除,


    | |   |   |   |   |   |   |     |   |
    ---- |----| ----|----| ---- | ---- |----|-----|-----|
    游标 | 2 | 5 | 6 | 4 | 0 | 3 | ... | 1 |
    数据 |   | A |   | D | E | B |     |   |
    下标 | 0 | 1 | 2 | 3 | 4 | 5 | ... | 999 | 


- ## <i id="4"></i>**优点**  
在插入和删除操作时,只需要修改游标,不需要移动元素,从而改进了在顺序存储结构中的插入和删除操作需要移动大量元素的缺点

- ## <i id="5"></i>**缺点**  
没有解决连续存储分配(数组)带来的表长难以确定的问题，失去了顺序存储结构随机存取的特征

