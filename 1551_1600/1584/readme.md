## [1584. 连接所有点的最小费用](https://leetcode-cn.com/problems/min-cost-to-connect-all-points/)

### 说明
给你一个points 数组，表示 2D 平面上的一些点，其中 points[i] = [xi, yi] 。

连接点 [xi, yi] 和点 [xj, yj] 的费用为它们之间的 曼哈顿距离 ：|xi - xj| + |yi - yj| ，其中 |val| 表示 val 的绝对值。

请你返回将所有点连接的最小总费用。只有任意两点之间 有且仅有 一条简单路径时，才认为所有点都已连接。

* 1 <= points.length <= 1000
* -106 <= xi, yi <= 106
* 所有点 (xi, yi) 两两不同。

### 实例
#### 1
输入：points = [[0,0],[2,2],[3,10],[5,2],[7,0]]
输出：20

#### 2
输入：points = [[3,12],[-2,5],[-4,1]]
输出：18