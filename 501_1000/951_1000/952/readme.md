## [952. 按公因数计算最大组件大小](https://leetcode.cn/problems/largest-component-size-by-common-factor/)

### 说明
给定一个由不同正整数的组成的非空数组 nums ，考虑下面的图：

有 nums.length 个节点，按从 nums[0] 到 nums[nums.length - 1] 标记；
只有当 nums[i] 和 nums[j] 共用一个大于 1 的公因数时，nums[i] 和 nums[j]之间才有一条边。
返回 图中最大连通组件的大小 。

### 提示：
* 1 <= nums.length <= 2 * 10^4
* 1 <= nums[i] <= 10^5
* nums 中所有值都 不同

### 实例
#### 1
输入：nums = [4,6,15,35]
输出：4

#### 2
输入：nums = [20,50,9,63]
输出：2