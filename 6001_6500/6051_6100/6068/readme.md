## [6068. 毯子覆盖的最多白色砖块数](https://leetcode.cn/problems/maximum-white-tiles-covered-by-a-carpet/)

### 说明
给你一个二维整数数组 tiles ，其中 tiles[i] = [li, ri] ，表示所有在 li <= j <= ri 之间的每个瓷砖位置 j 都被涂成了白色。

同时给你一个整数 carpetLen ，表示可以放在 任何位置 的一块毯子。

请你返回使用这块毯子，最多 可以盖住多少块瓷砖。

### 提示
* 1 <= tiles.length <= 5 * 104
* tiles[i].length == 2
* 1 <= li <= ri <= 109
* 1 <= carpetLen <= 109
* tiles 互相 不会重叠 。

### 实例
#### 1
输入：tiles = [[1,5],[10,11],[12,18],[20,25],[30,32]], carpetLen = 10
输出：9

#### 2
输入：tiles = [[10,11],[1,1]], carpetLen = 2
输出：2