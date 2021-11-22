## [2076. 处理含限制条件的好友请求](https://leetcode-cn.com/problems/process-restricted-friend-requests/)

### 说明
给你一个整数 n ，表示网络上的用户数目。每个用户按从 0 到 n - 1 进行编号。

给你一个下标从 0 开始的二维整数数组 restrictions ，
其中 restrictions[i] = [xi, yi] 意味着用户 xi 和用户 yi 不能 成为 朋友 ，不管是 直接 还是通过其他用户 间接 。
最初，用户里没有人是其他用户的朋友。给你一个下标从 0 开始的二维整数数组 requests 表示好友请求的列表，
其中 requests[j] = [uj, vj] 是用户 uj 和用户 vj 之间的一条好友请求。

如果 uj 和 vj 可以成为 朋友 ，那么好友请求将会 成功 。
每个好友请求都会按列表中给出的顺序进行处理（即，requests[j] 会在 requests[j + 1] 前）。
一旦请求成功，那么对所有未来的好友请求而言， uj 和 vj 将会 成为直接朋友 。

返回一个 布尔数组 result ，其中元素遵循此规则：如果第 j 个好友请求 成功 ，那么 result[j] 就是 true ；否则，为 false 。

注意：如果 uj 和 vj 已经是直接朋友，那么他们之间的请求将仍然 成功 。

* 2 <= n <= 1000
* 0 <= restrictions.length <= 1000
* restrictions[i].length == 2
* 0 <= xi, yi <= n - 1
* xi != yi
* 1 <= requests.length <= 1000
* requests[j].length == 2
* 0 <= uj, vj <= n - 1
* uj != vj

### 实例
#### 1
输入：n = 3, restrictions = [[0,1]], requests = [[0,2],[2,1]]
输出：[true,false]

#### 2
输入：n = 3, restrictions = [[0,1]], requests = [[1,2],[0,2]]
输出：[true,false]