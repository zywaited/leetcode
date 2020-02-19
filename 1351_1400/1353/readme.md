## [1353. 最多可以参加的会议数目](https://leetcode-cn.com/problems/maximum-number-of-events-that-can-be-attended/)

### 说明
给你一个数组 events，其中 events[i] = [startDayi, endDayi] ，表示会议 i 开始于 startDayi ，结束于 endDayi 。

你可以在满足 startDayi <= d <= endDayi 中的任意一天 d 参加会议 i 。注意，一天只能参加一个会议。

请你返回你可以参加的 最大 会议数目。

* 1 <= events.length <= 10^5
* events[i].length == 2
* 1 <= events[i][0] <= events[i][1] <= 10^5

### 实例
#### 1
输入：events = [[1,2],[2,3],[3,4]]
输出：3

#### 2
输入：events= [[1,2],[2,3],[3,4],[1,2]]
输出：4