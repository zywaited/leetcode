## [2057. 值相等的最小索引](https://leetcode.cn/problems/smallest-index-with-equal-value/)

### 说明
给你一个下标从 0 开始的整数数组 nums ，返回 nums 中满足 i mod 10 == nums[i] 的最小下标 i ；如果不存在这样的下标，返回 -1 。

x mod y 表示 x 除以 y 的 余数 。
请你返回一个整数数组 answer ，其中 answer[i] 是第 i 个查询的答案。

提示
* 1 <= nums.length <= 100
* 0 <= nums[i] <= 9

### 实例
#### 1
输入：nums = [0,1,2]
输出：0

#### 2
输入：nums = [4,3,2,1]
输出：2