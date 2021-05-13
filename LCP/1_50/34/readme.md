## [LCP 34. 二叉树染色](https://leetcode-cn.com/problems/er-cha-shu-ran-se-UGC/)

### 说明
小扣有一个根结点为 root 的二叉树模型，初始所有结点均为白色，可以用蓝色染料给模型结点染色，模型的每个结点有一个 val 价值。
小扣出于美观考虑，希望最后二叉树上每个蓝色相连部分的结点个数不能超过 k 个，求所有染成蓝色的结点价值总和最大是多少？

* 1 <= k <= 10
* 1 <= val <= 10000
* 1 <= 结点数量 <= 10000

### 实例
#### 1
输入：root = [5,2,3,4], k = 2
输出：12

#### 2
输入：root = [4,1,3,9,null,null,2], k = 2
输出：16