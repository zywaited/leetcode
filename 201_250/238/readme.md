## [除自身以外数组的乘积](https://leetcode-cn.com/problems/product-of-array-except-self/)
### 说明

给定长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。

### 要求
* 请不要使用除法，且在 O(n) 时间复杂度内完成此题。

### 实例
#### 1
输入: [1,2,3,4]
输出: [24,12,8,6]实现可以使用暴力法, 不过为了时间复杂度, 利用递减队列(也可是其他类型)保留当前的窗口最大值

### 实现
* one 空间复杂度O(N)
* twi 空间复杂度O(1)