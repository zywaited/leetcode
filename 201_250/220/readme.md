## [220. 存在重复元素 III](https://leetcode-cn.com/problems/contains-duplicate-iii/)

### 说明
给你一个整数数组 nums 和两个整数 k 和 t 。
请你判断是否存在 两个不同下标 i 和 j，使得 abs(nums[i] - nums[j]) <= t ，同时又满足 abs(i - j) <= k 。

如果存在则返回 true，不存在返回 false。

* 0 <= nums.length <= 2 * 10^4
* -2^31 <= nums[i] <= 2^31 - 1
* 0 <= k <= 10^4
* 0 <= t <= 2^31 - 1

### 实例
#### 1
输入：nums = [1,2,3,1], k = 3, t = 0
输出：true

#### 2
输入：nums = [1,0,1,1], k = 1, t = 2
输出：true