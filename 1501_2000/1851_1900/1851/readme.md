## [1851. 包含每个查询的最小区间](https://leetcode-cn.com/problems/minimum-interval-to-include-each-query/)

### 说明
给你一个二维整数数组 intervals ，其中 intervals[i] = [lefti, righti] 表示第 i 个区间开始于 lefti 、结束于 righti（包含两侧取值，闭区间）。
区间的 长度 定义为区间中包含的整数数目，更正式地表达是 righti - lefti + 1 。

再给你一个整数数组 queries 。第 j 个查询的答案是满足 lefti <= queries[j] <= righti 的 长度最小区间 i 的长度 。
如果不存在这样的区间，那么答案是 -1 。

以数组形式返回对应查询的所有答案。

提示
* 1 <= intervals.length <= 10^5
* 1 <= queries.length <= 10^5
* queries[i].length == 2
* 1 <= lefti <= righti <= 10^7
* 1 <= queries[j] <= 10^7

### 实例
#### 1
输入：intervals = [[1,4],[2,4],[3,6],[4,4]], queries = [2,3,4,5]
输出：[3,3,1,4]

#### 2
输入：intervals = [[2,3],[2,5],[1,8],[20,25]], queries = [2,19,5,22]
输出：[2,-1,4,6]