## [1380. 矩阵中的幸运数](https://leetcode-cn.com/problems/lucky-numbers-in-a-matrix/)

### 说明
给你一个 m * n 的矩阵，矩阵中的数字 各不相同 。请你按 任意 顺序返回矩阵中的所有幸运数。

幸运数是指矩阵中满足同时下列两个条件的元素：

在同一行的所有元素中最小
在同一列的所有元素中最大

* m == mat.length
* n == mat[i].length
* 1 <= n, m <= 50
* 1 <= matrix[i][j] <= 10^5
* 矩阵中的所有元素都是不同的

### 实例
#### 1
输入：matrix = [[3,7,8],[9,11,13],[15,16,17]]
输出：[15]

#### 2
输入：matrix = [[1,10,4,2],[9,3,8,7],[15,16,17,12]]
输出：[12]