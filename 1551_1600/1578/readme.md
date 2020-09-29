## [1578. 避免重复字母的最小删除成本](https://leetcode-cn.com/problems/minimum-deletion-cost-to-avoid-repeating-letters/)

### 说明
给你一个字符串 s 和一个整数数组 cost ，其中 cost[i] 是从 s 中删除字符 i 的代价。

返回使字符串任意相邻两个字母不相同的最小删除成本。

请注意，删除一个字符后，删除其他字符的成本不会改变。

* s.length == cost.length
* 1 <= s.length, cost.length <= 10^5
* 1 <= cost[i] <= 10^4
* s 中只含有小写英文字母

### 实例
#### 1
输入：s = "abaac", cost = [1,2,3,4,5]
输出：3

#### 2
输入：s = "abc", cost = [1,2,3]
输出：0