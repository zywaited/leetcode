## [2359. 找到离给定两个节点最近的节点](https://leetcode.cn/problems/find-closest-node-to-given-two-nodes/)

### 说明
给你一个 n 个节点的 有向图 ，节点编号为 0 到 n - 1 ，每个节点 至多 有一条出边。

有向图用大小为 n 下标从 0 开始的数组 edges 表示，表示节点 i 有一条有向边指向 edges[i] 。
如果节点 i 没有出边，那么 edges[i] == -1 。

同时给你两个节点 node1 和 node2 。

请你返回一个从 node1 和 node2 都能到达节点的编号，使节点 node1 和节点 node2 到这个节点的距离 较大值最小化。
如果有多个答案，请返回 最小 的节点编号。如果答案不存在，返回 -1 。

注意 edges 可能包含环。

### 提示：
* n == edges.length
* 2 <= n <= 10^5
* -1 <= edges[i] < n
* edges[i] != i
* 0 <= node1, node2 < n

### 实例
#### 1
输入：edges = [2,2,3,-1], node1 = 0, node2 = 1
输出：2

#### 2
输入：edges = [1,2,-1], node1 = 0, node2 = 2
输出：2