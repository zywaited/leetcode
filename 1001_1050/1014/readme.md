## [1014. 最佳观光组合](https://leetcode-cn.com/problems/best-sightseeing-pair/)

### 说明
给定正整数数组 A，A[i] 表示第 i 个观光景点的评分，并且两个景点 i 和 j 之间的距离为 j - i。

一对景点（i < j）组成的观光组合的得分为（A[i] + A[j] + i - j）：景点的评分之和减去它们两者之间的距离。

返回一对观光景点能取得的最高分。

* 2 <= A.length <= 50000
* 1 <= A[i] <= 1000

### 实例
#### 1
输入：[8,1,5,2,6]
输出：11