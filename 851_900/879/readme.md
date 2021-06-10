## [879. 盈利计划](https://leetcode-cn.com/problems/profitable-schemes/)

### 说明
集团里有 n 名员工，他们可以完成各种各样的工作创造利润。

第 i 种工作会产生 profit[i] 的利润，它要求 group[i] 名成员共同参与。如果成员参与了其中一项工作，就不能参与另一项工作。

工作的任何至少产生 minProfit 利润的子集称为 盈利计划 。并且工作的成员总数最多为 n 。

有多少种计划可以选择？因为答案很大，所以 返回结果模 10^9 + 7 的值。


* 1 <= n <= 100
* 0 <= minProfit <= 100
* 1 <= group.length <= 100
* 1 <= group[i] <= 100
* profit.length == group.length
* 0 <= profit[i] <= 100

### 实例
#### 1
输入：n = 5, minProfit = 3, group = [2,2], profit = [2,3]
输出：2

#### 2
输入：n = 10, minProfit = 5, group = [2,3,5], profit = [6,7,8]
输出：7