## [1383. 最大的团队表现值](https://leetcode-cn.com/problems/maximum-performance-of-a-team/)

### 说明
公司有编号为 1 到 n 的 n 个工程师，给你两个数组 speed 和 efficiency ，其中 speed[i] 和 efficiency[i] 分别代表第 i 位工程师的速度和效率。
请你返回由最多 k 个工程师组成的 ​​​​​​最大团队表现值 ，由于答案可能很大，请你返回结果对 10^9 + 7 取余后的结果。

团队表现值 的定义为：一个团队中「所有工程师速度的和」乘以他们「效率值中的最小值」。

* 1 <= n <= 10^5
* speed.length == n
* efficiency.length == n
* 1 <= speed[i] <= 10^5
* 1 <= efficiency[i] <= 10^8
* 1 <= k <= n

### 实例
#### 1
输入：n = 6, speed = [2,10,3,1,5,8], efficiency = [5,4,3,9,7,2], k = 2
输出：60

#### 2
输入：n = 6, speed = [2,10,3,1,5,8], efficiency = [5,4,3,9,7,2], k = 3
输出：68