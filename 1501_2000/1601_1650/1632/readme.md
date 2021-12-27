## [1632. 矩阵转换后的秩](https://leetcode-cn.com/problems/rank-transform-of-a-matrix/)

### 说明
给你一个 m x n 的矩阵 matrix ，请你返回一个新的矩阵 answer ，其中 answer[row][col] 是 matrix[row][col] 的秩。

每个元素的 秩 是一个整数，表示这个元素相对于其他元素的大小关系，它按照如下规则计算：

秩是从 1 开始的一个整数。
如果两个元素 p 和 q 在 同一行 或者 同一列 ，那么：
如果 p < q ，那么 rank(p) < rank(q)
如果 p == q ，那么 rank(p) == rank(q)
如果 p > q ，那么 rank(p) > rank(q)
秩 需要越 小 越好。
题目保证按照上面规则 answer 数组是唯一的。

* m == matrix.length
* n == matrix[i].length
* 1 <= m, n <= 500
* -10^9 <= matrix[row][col] <= 10^9

### 实例
#### 1
输入：matrix = [[1,2],[3,4]]
输出：[[1,2],[2,3]]

#### 2
输入：matrix = [[7,7],[7,7]]
输出：[[1,1],[1,1]]