## [5340. 统计有序矩阵中的负数](https://leetcode-cn.com/problems/count-negative-numbers-in-a-sorted-matrix/)

### 说明
给你一个 m * n 的矩阵 grid，矩阵中的元素无论是按行还是按列，都以非递增顺序排列。

请你统计并返回 grid 中 负数 的数目。

* m == grid.length
* n == grid[i].length
* 1 <= m, n <= 100
* -100 <= grid[i][j] <= 100

### 实例
#### 1
输入：grid = [[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]
输出：8
解释：矩阵中共有 8 个负数。

#### 2
输入：grid = [[3,2],[1,0]]
输出：0