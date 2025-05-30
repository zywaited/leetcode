## [3355. 零数组变换 I](https://leetcode.cn/problems/zero-array-transformation-i/)

### 说明
给定一个长度为 n 的整数数组 nums 和一个二维数组 queries，其中 queries[i] = [li, ri]。

对于每个查询 queries[i]：

在 nums 的下标范围 [li, ri] 内选择一个下标 子集。
将选中的每个下标对应的元素值减 1。
零数组 是指所有元素都等于 0 的数组。

如果在按顺序处理所有查询后，可以将 nums 转换为 零数组 ，则返回 true，否则返回 false。

### 提示：
* 1 <= nums.length <= 10^5
* 0 <= nums[i] <= 10^5
* 1 <= queries.length <= 10^5
* queries[i].length == 2
* 0 <= li <= ri < nums.length

### 实例
#### 1
输入：nums = [1,0,1], queries = [[0,2]]
输出：true

#### 2
输入：nums = [4,3,2,1], queries = [[1,3],[0,2]]
输出：false