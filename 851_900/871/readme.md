## [871. 最低加油次数](https://leetcode-cn.com/problems/minimum-number-of-refueling-stops/)

### 说明
汽车从起点出发驶向目的地，该目的地位于出发位置东面 target 英里处。

沿途有加油站，每个 station[i] 代表一个加油站，它位于出发位置东面 station[i][0] 英里处，并且有 station[i][1] 升汽油。

假设汽车油箱的容量是无限的，其中最初有 startFuel 升燃料。它每行驶 1 英里就会用掉 1 升汽油。

当汽车到达加油站时，它可能停下来加油，将所有汽油从加油站转移到汽车中。

为了到达目的地，汽车所必要的最低加油次数是多少？如果无法到达目的地，则返回 -1 。

注意：如果汽车到达加油站时剩余燃料为 0，它仍然可以在那里加油。如果汽车到达目的地时剩余燃料为 0，仍然认为它已经到达目的地。

* 1 <= target, startFuel, stations[i][1] <= 10^9
* 0 <= stations.length <= 500
* 0 < stations[0][0] < stations[1][0] < ... < stations[stations.length-1][0] < target

### 实例
#### 1
输入：target = 1, startFuel = 1, stations = []
输出：0

#### 2
输入：target = 100, startFuel = 1, stations = [[10,100]]
输出：-1