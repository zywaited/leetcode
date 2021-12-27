## [475. 供暖器](https://leetcode-cn.com/problems/heaters/)

### 说明
冬季已经来临。你的任务是设计一个有固定加热半径的供暖器向所有房屋供暖。

在加热器的加热半径范围内的每个房屋都可以获得供暖。

现在，给出位于一条水平线上的房屋houses 和供暖器heaters 的位置，请你找出并返回可以覆盖所有房屋的最小加热半径。

说明：所有供暖器都遵循你的半径标准，加热的半径也一样。

提示：
1 <= houses.length, heaters.length <= 3 * 10^4
1 <= houses[i], heaters[i] <= 10^9

### 实例
#### 1
输入: houses = [1,2,3], heaters = [2]
输出: 1

#### 2
输入: houses = [1,2,3,4], heaters = [1,4]
输出: 1