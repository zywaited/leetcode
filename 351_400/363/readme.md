## [363. 矩形区域不超过 K 的最大数值和](https://leetcode-cn.com/problems/max-sum-of-rectangle-no-larger-than-k/)

### 说明
给你一个 m x n 的矩阵 matrix 和一个整数 k ，找出并返回矩阵内部矩形区域的不超过 k 的最大数值和。

题目数据保证总会存在一个数值和不超过 k 的矩形区域。

* m == matrix.length
* n == matrix[i].length
* 1 <= m, n <= 100
* -100 <= matrix[i][j] <= 100
* -105 <= k <= 105

### 实例
#### 1
输入：matrix = [[1,0,1],[0,-2,3]], k = 2
输出：2

#### 2
输入：matrix = [[2,2,-1]], k = 3
输出：3