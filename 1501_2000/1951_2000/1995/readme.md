## [1995. 统计特殊四元组](https://leetcode-cn.com/problems/count-special-quadruplets/)

### 说明
给你一个 下标从 0 开始 的整数数组 nums ，返回满足下述条件的 不同 四元组 (a, b, c, d) 的 数目 ：
nums[a] + nums[b] + nums[c] == nums[d] ，且
a < b < c < d

提示
* 4 <= nums.length <= 50
* 1 <= nums[i] <= 100

### 实例
#### 1
输入：nums = [1,2,3,6]
输出：1

#### 2
输入：nums = [3,3,6,4,5]
输出：0