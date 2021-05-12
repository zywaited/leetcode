## [1310. 子数组异或查询](https://leetcode-cn.com/problems/xor-queries-of-a-subarray/)

### 说明
有一个正整数数组 arr，现给你一个对应的查询数组 queries，其中 queries[i] = [Li, Ri]。

对于每个查询 i，请你计算从 Li 到 Ri 的 XOR 值（即 arr[Li] xor arr[Li+1] xor ... xor arr[Ri]）作为本次查询的结果。

并返回一个包含给定查询 queries 所有结果的数组。

* 1 <= arr.length <= 3 * 10^4
* 1 <= arr[i] <= 10^9
* 1 <= queries.length <= 3 * 10^4
* queries[i].length == 2
* 0 <= queries[i][0] <= queries[i][1] < arr.length

### 实例
#### 1
输入：arr = [1,3,4,8], queries = [[0,1],[1,2],[0,3],[3,3]]
输出：[2,7,14,8]

#### 2
输入：arr = [4,8,2,10], queries = [[2,3],[1,3],[0,0],[0,3]]
输出：[8,0,4,4]