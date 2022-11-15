## [1770. 执行乘法运算的最大分数](https://leetcode.cn/problems/maximum-score-from-performing-multiplication-operations/)

### 说明
给你两个长度分别 n 和 m 的整数数组 nums 和 multipliers ，其中 n >= m ，数组下标 从 1 开始 计数。

初始时，你的分数为 0 。你需要执行恰好 m 步操作。在第 i 步操作（从 1 开始 计数）中，需要：

选择数组 nums 开头处或者末尾处 的整数 x 。
你获得 multipliers[i] * x 分，并累加到你的分数中。
将 x 从数组 nums 中移除。
在执行 m 步操作后，返回 最大 分数。

### 提示：
* n == nums.length
* m == multipliers.length
* 1 <= m <= 10^3
* m <= n <= 10^5
* -1000 <= nums[i], multipliers[i] <= 1000

### 实例
#### 1
输入：nums = [1,2,3], multipliers = [3,2,1]
输出：14

#### 2
输入：nums = [-5,-3,-3,-2,7,1], multipliers = [-10,-5,3,4,6]
输出：102