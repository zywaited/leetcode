## [2448. 使数组相等的最小开销](https://leetcode.cn/problems/minimum-cost-to-make-array-equal/)

### 说明
给你两个下标从 0 开始的数组 nums 和 cost ，分别包含 n 个 正 整数。

你可以执行下面操作 任意 次：

将 nums 中 任意 元素增加或者减小 1 。
对第 i 个元素执行一次操作的开销是 cost[i] 。

请你返回使 nums 中所有元素 相等 的 最少 总开销。

### 提示：
* n == nums.length == cost.length
* 1 <= n <= 10^5
* 1 <= nums[i], cost[i] <= 10^6

### 实例
#### 1
输入：nums = [1,3,5,2], cost = [2,3,1,14]
输出：8

#### 2
输入：nums = [2,2,2,2,2], cost = [4,2,8,1,3]
输出：0