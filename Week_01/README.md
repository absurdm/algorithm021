# 学习笔记

* 暂未详细整理整理

* 常见复杂度
  * O(1) < O(logn) < O(n) < O(nlogn) < O(n^2) < O(n^3) < O(2^n) < O(n!) < O(n^n)

## 数组 Array

* 数组操作:
  * 根据下标查询, 增加元素, 删除元素
* 根据下标查询: O(1)
* 插入一个元素:

* 时间复杂度
  * prepend O(n): 特殊优化后 O(1)
  * append O(1)
  * lookup O(1)
  * insert O(n)
  * delete O(n)

## 链表 Link List

* 时间复杂度
  * prepend O(1)
  * append O(1)
  * lookup O(n)
  * insert O(1)
  * delete O(1)

## 跳表 Skip List

  * 链表的特殊情况: 链表中的数据有序

### 给有序的链表优化查询

跳表: 升维思想+空间换时间

## 栈队列 Stack Queue

* 栈 Stack: 后入先出(先入后出) (LIFO, FILO)
* 队列 Queue: 先入先出(后入后出), (FIFO, LILO)

* 时间复杂度
  * prepend,append O(1)
  * lookup O(n)

### 双端队列 Deque: Double-End Queue

### 优先队列 Priority Queue
* 抽象的数据结构数据结构上有多重形成方式
1. insert O(1)
2. 取出 O(logn)

3. 底层具体实现的数据结构多样复杂
   1. 堆  heap
   2. 二叉搜索树 bst
   3. treap

### Golang Queue & Priority Queue 源码分析
