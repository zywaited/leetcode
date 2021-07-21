## [1829. 每个查询的最大异或值](https://leetcode-cn.com/problems/maximum-xor-for-each-query/)

### 说明
给你一个 有序 数组 nums ，它由 n 个非负整数组成，同时给你一个整数 maximumBit 。你需要执行以下查询 n 次：

找到一个非负整数 k < 2maximumBit ，使得 nums[0] XOR nums[1] XOR ... XOR nums[nums.length-1] XOR k 的结果 最大化 。k 是第 i 个查询的答案。
从当前数组 nums 删除 最后 一个元素。

请你返回一个数组 answer ，其中 answer[i]是第 i 个查询的结果。

* nums.length == n
* 1 <= n <= 105
* 1 <= maximumBit <= 20
* 0 <= nums[i] < 2maximumBit
* nums​​​ 中的数字已经按 升序 排好序。

### 实例
#### 1
输入：nums = [0,1,1,3], maximumBit = 2
输出：[0,3,2,3]

#### 2
输入：nums = [2,3,4,7], maximumBit = 3
输出：[5,2,6,5]