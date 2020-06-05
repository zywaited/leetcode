## [1409. 查询带键的排列](https://leetcode-cn.com/problems/queries-on-a-permutation-with-key/)

### 说明
给你一个待查数组 queries ，数组中的元素为 1 到 m 之间的正整数。 请你根据以下规则处理所有待查项 queries[i]（从 i=0 到 i=queries.length-1）：

一开始，排列 P=[1,2,3,...,m]。
对于当前的 i ，请你找出待查项 queries[i] 在排列 P 中的位置（下标从 0 开始），然后将其从原位置移动到排列 P 的起始位置（即下标为 0 处）。
注意， queries[i] 在 P 中的位置就是 queries[i] 的查询结果。
请你以数组形式返回待查数组  queries 的查询结果。

* 1 <= m <= 10^3
* 1 <= queries.length <= m
* 1 <= queries[i] <= m

### 实例
#### 1
输入：queries = [4,1,2,2], m = 4
输出：[3,1,2,0]

#### 2
输入：queries = [7,5,5,8,3], m = 8
输出：[6,5,0,7,5]