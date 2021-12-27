## [1395. 统计作战单位数](https://leetcode-cn.com/problems/count-number-of-teams/)

### 说明
n 名士兵站成一排。每个士兵都有一个 独一无二 的评分 rating 。

每 3 个士兵可以组成一个作战单位，分组规则如下：

从队伍中选出下标分别为 i、j、k 的 3 名士兵，他们的评分分别为 rating[i]、rating[j]、rating[k]
作战单位需满足： rating[i] < rating[j] < rating[k] 或者 rating[i] > rating[j] > rating[k] ，其中  0 <= i < j < k < n
请你返回按上述条件可以组建的作战单位数量。每个士兵都可以是多个作战单位的一部分。

* n == rating.length
* 1 <= n <= 200
* 1 <= rating[i] <= 10^5

### 实例
#### 1
输入：rating = [2,5,3,4,1]
输出：3

#### 2
输入：rating = [2,1,3]
输出：0