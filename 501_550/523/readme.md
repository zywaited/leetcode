## [523. 连续的子数组和](https://leetcode-cn.com/problems/continuous-subarray-sum/)

### 说明
给你一个整数数组 nums 和一个整数 k ，编写一个函数来判断该数组是否含有同时满足下述条件的连续子数组：

子数组大小 至少为 2 ，且
子数组元素总和为 k 的倍数。
如果存在，返回 true ；否则，返回 false 。

如果存在一个整数 n ，令整数 x 符合 x = n * k ，则称 x 是 k 的一个倍数。

* 1 <= nums.length <= 10^5
* 0 <= nums[i] <= 10^9
* 0 <= sum(nums[i]) <= 2^31 - 1
* 1 <= k <= 2^31 - 1

### 实例
#### 1
输入：nums = [23,2,4,6,7], k = 6
输出：true

#### 2
输入：nums = [23,2,6,4,7], k = 6
输出：true