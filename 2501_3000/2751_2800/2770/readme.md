## [2770. 达到末尾下标所需的最大跳跃次数](https://leetcode.cn/problems/maximum-number-of-jumps-to-reach-the-last-index/)

### 说明
给你一个下标从 0 开始、由 n 个整数组成的数组 nums 和一个整数 target 。

你的初始位置在下标 0 。在一步操作中，你可以从下标 i 跳跃到任意满足下述条件的下标 j ：
* 0 <= i < j < n
* -target <= nums[j] - nums[i] <= target
返回到达下标 n - 1 处所需的 最大跳跃次数 。

如果无法到达下标 n - 1 ，返回 -1 。

* 2 <= nums.length == n <= 1000
* -10^9 <= nums[i] <= 10^9
* 0 <= target <= 2 * 10^9

### 实例
#### 1
输入：nums = [1,3,6,4,1,2], target = 2
输出：3

#### 2
输入：nums = [1,3,6,4,1,2], target = 3
输出：5