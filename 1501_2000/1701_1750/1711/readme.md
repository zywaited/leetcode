## [1711. 大餐计数](https://leetcode-cn.com/problems/count-good-meals/)

### 说明
大餐 是指 恰好包含两道不同餐品 的一餐，其美味程度之和等于 2 的幂。

你可以搭配 任意 两道餐品做一顿大餐。

给你一个整数数组 deliciousness ，其中 deliciousness[i] 是第 i​​​​​​​​​​​​​​ 道餐品的美味程度，返回你可以用数组中的餐品做出的不同 大餐 的数量。
结果需要对 10^9 + 7 取余。

注意，只要餐品下标不同，就可以认为是不同的餐品，即便它们的美味程度相同。

* 1 <= deliciousness.length <= 10^5
* 0 <= deliciousness[i] <= 2^20

### 实例
#### 1
输入：deliciousness = [1,3,5,7,9]
输出：4

#### 2
输入：deliciousness = [1,1,1,3,3,3,7]
输出：15