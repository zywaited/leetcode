## [6164. 数位和相等数对的最大和](https://leetcode.cn/problems/max-sum-of-a-pair-with-equal-sum-of-digits/)

### 说明
给你一个下标从 0 开始的数组 nums ，数组中的元素都是 正 整数。
请你选出两个下标 i 和 j（i != j），且 nums[i] 的数位和 与  nums[j] 的数位和相等。

请你找出所有满足条件的下标 i 和 j ，找出并返回 nums[i] + nums[j] 可以得到的 最大值 。

### 提示：
* 1 <= nums.length <= 10^5
* 1 <= nums[i] <= 10^9

### 实例
#### 1
输入：nums = [18,43,36,13,7]
输出：54

#### 2
输入：nums = [10,12,19,14]
输出：-1