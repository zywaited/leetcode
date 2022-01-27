## [2050. 并行课程 III](https://leetcode-cn.com/problems/parallel-courses-iii/)

### 说明
给你一个整数 n ，表示有 n 节课，课程编号从 1 到 n 。
同时给你一个二维整数数组 relations ，其中 relations[j] = [prevCoursej, nextCoursej] ，表示课程 prevCoursej 必须在课程 nextCoursej 之前 完成（先修课的关系）。
同时给你一个下标从 0 开始的整数数组 time ，其中 time[i] 表示完成第 (i+1) 门课程需要花费的 月份 数。

请你根据以下规则算出完成所有课程所需要的 最少 月份数：

如果一门课的所有先修课都已经完成，你可以在 任意 时间开始这门课程。
你可以 同时 上 任意门课程 。
请你返回完成所有课程所需要的 最少 月份数。

注意：测试数据保证一定可以完成所有课程（也就是先修课的关系构成一个有向无环图）。

提示
* 1 <= n <= 5 * 10^4
* 0 <= relations.length <= min(n * (n - 1) / 2, 5 * 10^4)
* relations[j].length == 2
* 1 <= prevCoursej, nextCoursej <= n
* prevCoursej != nextCoursej
* 所有的先修课程对 [prevCoursej, nextCoursej] 都是 互不相同 的。
* time.length == n
* 1 <= time[i] <= 10^4
* 先修课程图是一个有向无环图。

### 实例
#### 1
输入：n = 3, relations = [[1,3],[2,3]], time = [3,2,5]
输出：8

#### 2
输入：n = 5, relations = [[1,5],[2,5],[3,5],[3,4],[4,5]], time = [1,2,3,4,5]
输出：12