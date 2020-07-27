## [1027. 最长等差数列](https://leetcode-cn.com/problems/longest-arithmetic-sequence/)

### 说明
给定一个整数数组 A，返回 A 中最长等差子序列的长度。
回想一下，A 的子序列是列表 A[i_1], A[i_2], ..., A[i_k] 其中 0 <= i_1 < i_2 < ... < i_k <= A.length - 1。
并且如果 B[i+1] - B[i]( 0 <= i < B.length - 1) 的值都相同，那么序列 B 是等差的。

* 2 <= A.length <= 2000
* 0 <= A[i] <= 10000

### 实例
#### 1
输入：[3,6,9,12]
输出：4

#### 2
输入：[9,4,7,2,10]
输出：3

### 实现
* one 动态规划 (数组+HASHMAP)