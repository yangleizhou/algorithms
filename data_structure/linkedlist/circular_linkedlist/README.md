## 循环链表 circular linkedlist

- [单向链表](../single_linkedlist)
- [双向链表](../double_linkedlist)
- [定义](#1)
- [分类](#2)


- ## <i id="1"></i>**`定义`**  
`循环链表`是另一种形式的链式存贮结构。它的特点是表中最后一个结点的`指针域`指向头结点，整个链表形成一个环。

- ## <i id="2"></i>**`分类`**  
    - [单向循环链表](./single_circular_linkedlist)
        1. 对于[单向链表](../single_linkedlist)中任何结点的访问都需要从`头结点`开始;而`单向循环链表`可以从任何一个结点开始,顺序向后访问到达任意结点。
    - [双向循环链表](./double_circular_linkedlist)
        2. 对于[双向链表](../double_linkedlist)中任意结点的访问只能从`头结点`或`尾结点`开始;而`双向循环链表`可以从任何结点开始任意向前向后双向访问。
