# 1. 题目描述

```
Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right, then right to left for the next level and alternate between).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its zigzag level order traversal as:
[
  [3],
  [20,9],
  [15,7]
]

```

# 2. 简单分析

相比前一题 `binary tree level order traversal`, 输出的要求中新增了 `zigzag`, 层序遍历的结果需要以锯齿形的方式输出

即单数层级的所有节点从左到右输出, 双数层级的所有节点从右向左输出, 在我们借助层序规律, 由当前出队元素层级+1 的方式获取其层级的基础上去判断是否需要将当前数组反转即可

换言之, 此题相比上一题, 新引入的要求是判断每层的数组是否需要反转, 此外, 还隐藏着一个坑:

> 最后一层由于需要做单独处理, 因此最后一层需要进行同样的判断来决定是否需要反转

新逻辑:

```go
///~ 循环出队已结束
if len(buf) != 0 {
    // if there is sth(last level) in buf
    if currLevel % 2 == 0 {
        rst = append(rst, buf)
    }
}
return rst
```
