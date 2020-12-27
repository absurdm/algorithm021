# 第四周

1. [第四周](#第四周)
   1. [深度优先搜索 Deep First Search DFS](#深度优先搜索-deep-first-search-dfs)
   2. [广度优先搜索 Breath First Search BFS](#广度优先搜索-breath-first-search-bfs)
   3. [贪心算法 Greedy](#贪心算法-greedy)
   4. [二分查找](#二分查找)
      1. [牛顿迭代法](#牛顿迭代法)


## 深度优先搜索 Deep First Search DFS

```python
# 递归写法 - 栈
visited = set()
def dfs(node, vistided):
    visited.addd(node)
    #  process current node here
    ...
    for  next_node in node.children:
        # 递归结束
        if not next_node in visited:
            # 下潜到下一层
            dfs(next node, visited)

# 非递归写法 - 手动维护栈

def dfs(self, tree):
    if tree.root is None:
        return []

    visited, stack =  [], [tree.root]

    while stack:
        node = stack.pop()
        visited.add(node)

        process(node)
        nodes = generate_related_nodes(node)
        stack.push(nodes)

    # other processing work
```

## 广度优先搜索 Breath First Search BFS

```python
# 广度优先遍历 - 队列
def BFS(gragh, start, end):
    visited =  set()
    queue = []
    queue.append([start])

    while queue:
        node = queue.pop()
        visited.add(node)

        process(node)
        nodes = generate_related_nodes(node)
        queue.push(nodes)
```

## 贪心算法 Greedy

每一步选择都采取当前状态下最好或者最优(最有利)的选择, 从而希望导致结果是全局最好或者最优的算法.

* 常用解决一些最优化的问题:
  * 最小生成树, 求哈弗曼编码.
  * 问题可以用贪心法来解决, 那么贪心算法一般是解决这个问题的一种最好办法.贪心法的高效性可以一级所求的答案比较接近最优结果, 贪心发可以辅助算法或者直接解决一些要求结果不是特别准确的问题.


* 适用贪心算法的场景:
  * 问题能否用贪心法, 需要根据问题实际情况来分析
  * 问题能够分解成子问题来解决, 子问题的最优解能递推到最终问题的最优解.这种子问题最优解构为最优子结构

* 使用贪心算法正反例:
  * 正例: 找零钱 20, 10, 5 => 个零钱存在整除关系 => 大问题能够用子问题递推得到结果 => 有最优子结构
  * 反例: 找零钱 10, 9, 1 => 个零钱不能完全整除 => 大问题不能用子问题递推得到结果 => 无最优子结构

* 贪心和动态规划的不同
  * 贪心算法与动态规划的不同在于它对每个子问题的解决方案都作出选择, 不能回退.动态规划则会保存以前的运算结果, 并根据以前的结果对当前进行选择,有回退功能.


## 二分查找

* 三个前提条件
  * 目标函数单调性(单调递增或者单调递减)
  * 存在上下界 (最大值, 最小值, Bounded)
  * 能够通过索引访问(数据组织结构, Index  accessible)

* 根据单调性判断排除掉前半部分或者后半部分

```python
left, right = 0, len(array)-1
while left <= right:
    mid = (left + right) / 2
    if array[mid] == target:
        break or return result
    elif array[mid] < target:
        left = mid + 1
    else:
        right = mid - 1

// 返回失败结果
return false
```

* LeetCode 33 题总结
    * left, right 为指针, 要区分指针和具体值关系
    * **二分查找变异问题: 断定在某个半区的条件即可, 另外半区 else 划过**


### 牛顿迭代法

* 暂未总结