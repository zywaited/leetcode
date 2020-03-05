## [994. 腐烂的橘子](https://leetcode-cn.com/problems/rotting-oranges/)

### 说明
在给定的网格中，每个单元格可以有以下三个值之一：

值 0 代表空单元格；
值 1 代表新鲜橘子；
值 2 代表腐烂的橘子。
每分钟，任何与腐烂的橘子（在 4 个正方向上）相邻的新鲜橘子都会腐烂。

返回直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1。

* 1 <= grid.length <= 10
* 1 <= grid[0].length <= 10
* grid[i][j] 仅为 0、1 或 2

### 实例
#### 1
输入：[[2,1,1],[1,1,0],[0,1,1]]
输出：4

#### 2
输入：[[2,1,1],[0,1,1],[1,0,1]]
输出：-1