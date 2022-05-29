## [6080. 使数组按非递减顺序排列](https://leetcode.cn/problems/steps-to-make-array-non-decreasing/)

### 说明
给你一个下标从 0 开始的整数数组 nums 。在一步操作中，移除所有满足 nums[i - 1] > nums[i] 的 nums[i] ，其中 0 < i < nums.length 。

重复执行步骤，直到 nums 变为 非递减 数组，返回所需执行的操作数。

### 提示：
* 1 <= nums.length <= 105
* 1 <= nums[i] <= 109

### 实例
#### 1
输入：nums = [5,3,4,4,7,3,6,11,8,5,11]
输出：3

#### 2
输入：nums = [4,5,7,7,13]
输出：0