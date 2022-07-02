## [1359. 有效的快递序列数目](https://leetcode.cn/problems/count-all-valid-pickup-and-delivery-options/)

### 说明
给你 n 笔订单，每笔订单都需要快递服务。

请你统计所有有效的 收件/配送 序列的数目，确保第 i 个物品的配送服务 delivery(i) 总是在其收件服务 pickup(i) 之后。

由于答案可能很大，请返回答案对 10^9 + 7 取余的结果。

### 提示：
* 1 <= n <= 500

### 实例
#### 1
输入：n = 1
输出：1

#### 2
输入：n = 2
输出：6