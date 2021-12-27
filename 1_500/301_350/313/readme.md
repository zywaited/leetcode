## [313. 超级丑数](https://leetcode-cn.com/problems/super-ugly-number/)

### 说明
编写一段程序来查找第 n 个超级丑数。

超级丑数是指其所有质因数都是长度为 k 的质数列表 primes 中的正整数。

* 1 是任何给定 primes 的超级丑数。
* 给定 primes 中的数字以升序排列。
* 0 < k ≤ 100, 0 < n ≤ 106, 0 < primes[i] < 1000 。
* 第 n 个超级丑数确保在 32 位有符整数范围内。

### 实例
#### 1
输入: n = 12, primes = [2,7,13,19]
输出: 32