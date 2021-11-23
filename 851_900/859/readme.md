## [859. 亲密字符串](https://leetcode-cn.com/problems/buddy-strings/)

### 说明
给你两个字符串 s 和 goal ，只要我们可以通过交换 s 中的两个字母得到与 goal 相等的结果，就返回 true ；否则返回 false 。

交换字母的定义是：取两个下标 i 和 j （下标从 0 开始）且满足 i != j ，接着交换 s[i] 和 s[j] 处的字符。

例如，在 "abcd" 中交换下标 0 和下标 2 的元素可以生成 "cbad" 。

* 1 <= s.length, goal.length <= 2 * 10^4
* s 和 goal 由小写英文字母组成

### 实例
#### 1
输入：s = "ab", goal = "ba"
输出：true

#### 2
输入：s = "ab", goal = "ab"
输出：false