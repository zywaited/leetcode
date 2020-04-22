## [面试题 08.14. 布尔运算](https://leetcode-cn.com/problems/boolean-evaluation-lcci/)

### 说明
给定一个布尔表达式和一个期望的布尔结果 result，
布尔表达式由 0 (false)、1 (true)、& (AND)、 | (OR) 和 ^ (XOR) 符号组成。
实现一个函数，算出有几种可使该表达式得出 result 值的括号方法。

* 运算符的数量不超过 19 个

### 实例
#### 1
输入: s = "1^0|0|1", result = 0
输出: 2

#### 2
输入: s = "0&0&0&1^1|0", result = 1
输出: 10