## [1642. 可以到达的最远建筑](https://leetcode-cn.com/problems/furthest-building-you-can-reach/)

### 说明
给你一个整数数组 heights ，表示建筑物的高度。另有一些砖块 bricks 和梯子 ladders 。

你从建筑物 0 开始旅程，不断向后面的建筑物移动，期间可能会用到砖块或梯子。

当从建筑物 i 移动到建筑物 i+1（下标 从 0 开始）时：

如果当前建筑物的高度 大于或等于 下一建筑物的高度，则不需要梯子或砖块
如果当前建筑的高度 小于 下一个建筑的高度，您可以使用 一架梯子 或 (h[i+1] - h[i]) 个砖块
如果以最佳方式使用给定的梯子和砖块，返回你可以到达的最远建筑物的下标（下标 从 0 开始 ）。

* 1 <= heights.length <= 10^5
* 1 <= heights[i] <= 10^6
* 0 <= bricks <= 10^9
* 0 <= ladders <= heights.length

### 实例
#### 1
输入：heights = [4,2,7,6,9,14,12], bricks = 5, ladders = 1
输出：4

#### 2
输入：heights = [4,12,2,7,3,18,20,3,19], bricks = 10, ladders = 2
输出：7