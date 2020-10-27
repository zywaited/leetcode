## [1365. 有多少小于当前数字的数字](https://leetcode-cn.com/problems/how-many-numbers-are-smaller-than-the-current-number/)

### 说明
给你一个数组 nums，对于其中每个元素 nums[i]，请你统计数组中比它小的所有数字的数目。

换而言之，对于每个 nums[i] 你必须计算出有效的 j 的数量，其中 j 满足 j != i 且 nums[j] < nums[i] 。

以数组形式返回答案。

* 2 <= nums.length <= 500
* 0 <= nums[i] <= 100

### 实例
#### 1
输入：nums = [8,1,2,2,3]
输出：[4,0,1,1,3]

#### 2
输入：nums = [6,5,4,8]
输出：[2,1,0,3]