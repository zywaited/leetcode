## [两数相除](https://leetcode-cn.com/problems/divide-two-integers/)
### 说明

给定两个整数，被除数 dividend 和除数 divisor。

返回被除数 dividend 除以除数 divisor 得到的商。

被除数和除数均为 32 位有符号整数。
除数不为 0。
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。

### 实例
#### 1

输入: dividend = 10, divisor = 3
输出: 3

#### 2

输入: dividend = 7, divisor = -3
输出: -2

### 要求

* 不使用乘法、除法和 mod 运算符

### 实现
* one 减法，不过该题有时间限制，所以减法需要做优化，需要做增长减法