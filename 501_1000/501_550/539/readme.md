## [539. 最小时间差](https://leetcode-cn.com/problems/minimum-time-difference/)

### 说明
给定一个 24 小时制（小时:分钟 "HH:MM"）的时间列表，找出列表中任意两个时间的最小时间差并以分钟数表示。

提示
* 2 <= timePoints <= 2 * 10^4
* timePoints[i] 格式为 "HH:MM"

### 实例
#### 1
输入：timePoints = ["23:59","00:00"]
输出：1

#### 2
输入：timePoints = ["00:00","23:59","00:00"]
输出：0

### 实现方法
* 基于排序
* 基于红黑树
* 基于数组