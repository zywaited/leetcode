## [1871. 跳跃游戏 VII](https://leetcode-cn.com/problems/jump-game-vii/)

### 说明
给你一个下标从 0 开始的二进制字符串 s 和两个整数 minJump 和 maxJump 。
一开始，你在下标 0 处，且该位置的值一定为 '0' 。
当同时满足如下条件时，你可以从下标 i 移动到下标 j 处：

i + minJump <= j <= min(i + maxJump, s.length - 1) 且
s[j] == '0'.
如果你可以到达 s 的下标 s.length - 1 处，请你返回 true ，否则返回 false 。

*  2 <= s.length <= 10^5
* s[i] 要么是 '0' ，要么是 '1'
* s[0] == '0'
* 1 <= minJump <= maxJump < s.length

### 实例
#### 1
输入：s = "011010", minJump = 2, maxJump = 3
输出：true

#### 2
输入：s = "01101110", minJump = 2, maxJump = 3
输出：false