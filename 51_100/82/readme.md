## [删除排序链表中的重复元素 II](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/)
### 说明

给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。

### 实例
#### 1

输入: 1->2->3->3->4->4->5
输出: 1->2->5

#### 2

输入: 1->1->1->2->3
输出: 2->3

### 实现
* one 就在原链表上改动
* two hashmap(自行实现)