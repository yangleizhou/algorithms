## 静态链表 static list

- [定义](#1)
- [实现原理](#2)
- [优点](#3)
- [缺点](#4)

https://www.cnblogs.com/dongry/p/10210609.html
https://blog.csdn.net/qq_23856059/article/details/81392649
https://www.jianshu.com/p/0f05c2fb3d81
https://blog.csdn.net/hhhhhyyyyy8/article/details/81027728
https://blog.csdn.net/tangguangqiang/article/details/53614789


- ## <i id="1"></i>**`定义`**  
`静态链表`,用[数组](../../array)来实现链式存储结构。兼顾`顺序表`和链表的优点，是`顺序表`和`链表`的升级,`静态链表`的数据全部存储在数组中(顺序表)，但存储的位置是随机的，数据直接的一对一关系是通过`游标`维持。</br>
 结点：数据域,用于存储数据元素的值。  
>>  游标, 即数组下标，表示直接后继元素所在数组中的位置。  

- ## <i id="2"></i>**`实现原理`**  
1. 使用结构体数组,结构体有游标cursor 和数据域 data  
2. 一个数组分量表示一个结点，用cursor代替指针指示结点在数组中的相对位置  
3. 数组逻辑上分为两个链表：`备用链表`(空闲的结点)和`数据链表`(已被使用的结点)