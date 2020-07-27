## [1026. 节点与其祖先之间的最大差值](https://leetcode-cn.com/problems/maximum-difference-between-node-and-ancestor/)

### 说明
给定二叉树的根节点 root，找出存在于不同节点 A 和 B 之间的最大值 V，其中 V = |A.val - B.val|，且 A 是 B 的祖先。

（如果 A 的任何子节点之一为 B，或者 A 的任何子节点是 B 的祖先，那么我们认为 A 是 B 的祖先）

* 树中的节点数在 2 到 5000 之间。
* 每个节点的值介于 0 到 100000 之间。

### 实例
#### 1
输入：[8,3,10,1,6,null,14,null,null,4,7,13]
输出：7