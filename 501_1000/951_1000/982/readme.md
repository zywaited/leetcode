## [982. 按位与为零的三元组](https://leetcode-cn.com/problems/triples-with-bitwise-and-equal-to-zero/)

### 说明
给定一个整数数组 A，找出索引为 (i, j, k) 的三元组，使得：

0 <= i < A.length
0 <= j < A.length
0 <= k < A.length
A[i] & A[j] & A[k] == 0，其中 & 表示按位与（AND）操作符。

* 1 <= A.length <= 1000
* 0 <= A[i] < 2^16

### 实例
#### 1
输入：[2,1,3]
输出：12