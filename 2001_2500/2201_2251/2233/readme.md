## [2233. K 次增加后的最大乘积](https://leetcode-cn.com/problems/maximum-product-after-k-increments/)

### 说明
给你一个非负整数数组 nums 和一个整数 k 。每次操作，你可以选择 nums 中 任一 元素并将它 增加 1 。

请你返回 至多 k 次操作后，能得到的 nums的 最大乘积 。由于答案可能很大，请你将答案对 109 + 7 取余后返回。

提示
* 1 <= nums.length, k <= 10^5
* 0 <= nums[i] <= 10^6

### 实例
#### 1
输入：nums = [0,4], k = 5
输出：20

#### 2
输入：nums = [6,3,3,2], k = 2
输出：216