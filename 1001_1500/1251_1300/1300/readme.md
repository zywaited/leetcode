## [1300. 转变数组后最接近目标值的数组和](https://leetcode-cn.com/problems/sum-of-mutated-array-closest-to-target/)

### 说明
给你一个整数数组 arr 和一个目标值 target ，请你返回一个整数 value ，
使得将数组中所有大于 value 的值变成 value 后，数组的和最接近  target （最接近表示两者之差的绝对值最小）。

如果有多种使得和最接近 target 的方案，请你返回这些整数中的最小值。

请注意，答案不一定是 arr 中的数字。

* 1 <= arr.length <= 10^4
* 1 <= arr[i], target <= 10^5

### 实例
#### 1
输入：arr = [4,9,3], target = 10
输出：3

#### 2
输入：arr = [2,3,5], target = 10
输出：5