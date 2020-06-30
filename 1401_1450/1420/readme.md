## [1420. 生成数组](https://leetcode-cn.com/problems/build-array-where-you-can-find-the-maximum-exactly-k-comparisons/)

### 说明
给你三个整数 n、m 和 k 。下图描述的算法用于找出正整数数组中最大的元素。
请你生成一个具有下述属性的数组 arr ：

arr 中有 n 个整数。
1 <= arr[i] <= m 其中 (0 <= i < n) 。
将上面提到的算法应用于 arr ，search_cost 的值等于 k 。
返回上述条件下生成数组 arr 的 方法数 ，由于答案可能会很大，所以 必须 对 10^9 + 7 取余。

* 1 <= n <= 50
* 1 <= m <= 100
* 0 <= k <= n

### 实例
#### 1
输入：n = 2, m = 3, k = 1
输出：6

#### 2
输入：n = 5, m = 2, k = 3
输出：0