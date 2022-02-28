## [寻找重复数](https://leetcode-cn.com/problems/find-the-duplicate-number/)

### 说明

给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。
数组中只有一个重复的数字，但它可能不止重复出现一次

### 实例
#### 1

输入: [1,3,4,2,2]
输出: 2

### 要求

* 不能更改原数组（假设数组是只读的）
* 只能使用额外的 O(1) 的空间
* 时间复杂度小于 O(n2) 

### 实现
如果没有1，2点要求，可使用的方法
1 排序;
2 hash map;
3 1-n转换为索引，对应位置为负数，如果这个位已经为负数，这个索引就是重复值;
* 快慢指针
* 二分查找法(重复值会导致另一边失衡)