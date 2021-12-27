## [327. 区间和的个数](https://leetcode-cn.com/problems/count-of-range-sum/)

### 说明
给你一个整数数组 nums 以及两个整数 lower 和 upper 。
求数组中，值位于范围 [lower, upper] （包含 lower 和 upper）之内的 区间和的个数 。

区间和 S(i, j) 表示在 nums 中，位置从 i 到 j 的元素之和，包含 i 和 j (i ≤ j)。

* 1 <= nums.length <= 105
* -2^31 <= nums[i] <= 2^31 - 1
* -10^5 <= lower <= upper <= 10^5
* 题目数据保证答案是一个 32 位 的整数

### 实例
#### 1
输入：nums = [-2,5,-1], lower = -2, upper = 2
输出：3

#### 2
输入：nums = [0], lower = 0, upper = 0
输出：1