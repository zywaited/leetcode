## [473. 火柴拼正方形](https://leetcode.cn/problems/matchsticks-to-square/)

### 说明
你将得到一个整数数组 matchsticks ，其中 matchsticks[i] 是第 i 个火柴棒的长度。
你要用 所有的火柴棍 拼成一个正方形。你 不能折断 任何一根火柴棒，但你可以把它们连在一起，而且每根火柴棒必须 使用一次 。

如果你能使这个正方形，则返回 true ，否则返回 false 。

### 提示
* 1 <= matchsticks.length <= 15
* 1 <= matchsticks[i] <= 108

### 实例
#### 1
输入: matchsticks = [1,1,2,2,2]
输出: true

#### 2
输入: matchsticks = [3,3,3,3,4]
输出: false