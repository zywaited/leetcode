## [2391. 收集垃圾的最少总时间](https://leetcode.cn/problems/minimum-amount-of-time-to-collect-garbage/)

### 说明
给你一个下标从 0 开始的字符串数组 garbage ，其中 garbage[i] 表示第 i 个房子的垃圾集合。garbage[i] 只包含字符 'M' ，'P' 和 'G' ，但可能包含多个相同字符，每个字符分别表示一单位的金属、纸和玻璃。垃圾车收拾 一 单位的任何一种垃圾都需要花费 1 分钟。

同时给你一个下标从 0 开始的整数数组 travel ，其中 travel[i] 是垃圾车从房子 i 行驶到房子 i + 1 需要的分钟数。

城市里总共有三辆垃圾车，分别收拾三种垃圾。每辆垃圾车都从房子 0 出发，按顺序 到达每一栋房子。但它们 不是必须 到达所有的房子。

任何时刻只有 一辆 垃圾车处在使用状态。当一辆垃圾车在行驶或者收拾垃圾的时候，另外两辆车 不能 做任何事情。

请你返回收拾完所有垃圾需要花费的 最少 总分钟数。

### 提示：
* 2 <= garbage.length <= 10^5
* garbage[i] 只包含字母 'M' ，'P' 和 'G' 。
* 1 <= garbage[i].length <= 10
* travel.length == garbage.length - 1
* 1 <= travel[i] <= 100

### 实例
#### 1
输入：garbage = ["G","P","GP","GG"], travel = [2,4,3]
输出：21

#### 2
输入：garbage = ["MMM","PGM","GP"], travel = [3,10]
输出：37