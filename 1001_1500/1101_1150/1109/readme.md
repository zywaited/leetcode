## [1109. 航班预订统计](https://leetcode-cn.com/problems/corporate-flight-bookings/)

### 说明
这里有 n 个航班，它们分别从 1 到 n 进行编号。

有一份航班预订表 bookings ，表中第 i 条预订记录 bookings[i] = [firsti, lasti, seatsi] 意味着在从 fi
rsti 到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi 个座位。

请你返回一个长度为 n 的数组 answer，其中 answer[i] 是航班 i 上预订的座位总数。

* 1 <= n <= 2 * 10^4
* 1 <= bookings.length <= 2 * 10^4
* bookings[i].length == 3
* 1 <= firsti <= lasti <= n
* 1 <= seatsi <= 10^4

### 实例
#### 1
输入：bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
输出：[10,55,45,25,25]

#### 2
输入：bookings = [[1,2,10],[2,2,15]], n = 2
输出：[10,25]