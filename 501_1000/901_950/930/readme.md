## [930. 和相同的二元子数组](https://leetcode-cn.com/problems/binary-subarrays-with-sum/)

### 说明
给你一个二元数组 nums ，和一个整数 goal ，请你统计并返回有多少个和为 goal 的 非空 子数组。

子数组 是数组的一段连续部分。

* 1 <= nums.length <= 3 * 104
* nums[i] 不是 0 就是 1
* 0 <= goal <= nums.length

### 实例
#### 1
输入：nums = [1,0,1,0,1], goal = 2
输出：4

#### 2
输入：nums = [0,0,0,0,0], goal = 0
输出：15