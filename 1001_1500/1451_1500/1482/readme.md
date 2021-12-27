## [1482. 制作 m 束花所需的最少天数](https://leetcode-cn.com/problems/minimum-number-of-days-to-make-m-bouquets/)

### 说明
给你一个整数数组 bloomDay，以及两个整数 m 和 k 。

现需要制作 m 束花。制作花束时，需要使用花园中 相邻的 k 朵花 。

花园中有 n 朵花，第 i 朵花会在 bloomDay[i] 时盛开，恰好 可以用于 一束 花中。

请你返回从花园中摘 m 束花需要等待的最少的天数。如果不能摘到 m 束花则返回 -1 。

* bloomDay.length == n
* 1 <= n <= 10^5
* 1 <= bloomDay[i] <= 10^9
* 1 <= m <= 10^6
* 1 <= k <= n

### 实例
#### 1
输入：bloomDay = [1,10,3,10,2], m = 3, k = 1
输出：3

#### 2
输入：bloomDay = [1,10,3,10,2], m = 3, k = 2
输出：-1