## [2834. 找出美丽数组的最小和](https://leetcode.cn/problems/find-the-minimum-possible-sum-of-a-beautiful-array/)

### 说明
给你两个正整数：n 和 target 。

如果数组 nums 满足下述条件，则称其为 美丽数组 。

nums.length == n.
nums 由两两互不相同的正整数组成。
在范围 [0, n-1] 内，不存在 两个 不同 下标 i 和 j ，使得 nums[i] + nums[j] == target 。
返回符合条件的美丽数组所可能具备的 最小 和，并对结果进行取模 10^9 + 7。

* 1 <= n <= 10^9
* 1 <= target <= 10^9

### 实例
#### 1
输入：n = 2, target = 3
输出：4

#### 2
输入：n = 3, target = 3
输出：8