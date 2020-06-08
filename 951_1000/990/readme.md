## [990. 等式方程的可满足性](https://leetcode-cn.com/problems/satisfiability-of-equality-equations/)

### 说明
给定一个由表示变量之间关系的字符串方程组成的数组，每个字符串方程 equations[i] 的长度为 4，
并采用两种不同的形式之一："a==b" 或 "a!=b"。在这里，a 和 b 是小写字母（不一定不同），表示单字母变量名。

只有当可以将整数分配给变量名，以便满足所有给定的方程时才返回 true，否则返回 false。

* 1 <= equations.length <= 500
* equations[i].length == 4
* equations[i][0] 和 equations[i][3] 是小写字母
* equations[i][1] 要么是 '='，要么是 '!'
* equations[i][2] 是 '='

### 实例
#### 1
输入：["a==b","b!=a"]
输出：false

#### 2
输出：["b==a","a==b"]
输入：true