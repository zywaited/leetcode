## [805. 数组的均值分割](https://leetcode.cn/problems/split-array-with-same-average/)

### 说明
给定你一个整数数组 nums

我们要将 nums 数组中的每个元素移动到 A 数组 或者 B 数组中，使得 A 数组和 B 数组不为空，并且 average(A) == average(B) 。

如果可以完成则返回true ， 否则返回 false  。

注意：对于数组 arr ,  average(arr) 是 arr 的所有元素的和除以 arr 长度。

* 1 <= nums.length <= 30
* 0 <= nums[i] <= 10^4

### 实例
#### 1
输入: nums = [1,2,3,4,5,6,7,8]
输出: true