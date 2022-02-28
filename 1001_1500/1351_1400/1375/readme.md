## [1375. 灯泡开关 III](https://leetcode-cn.com/problems/bulb-switcher-iii/)

### 说明
房间中有 n 枚灯泡，编号从 1 到 n，自左向右排成一排。最初，所有的灯都是关着的。
在 k  时刻（ k 的取值范围是 0 到 n - 1），我们打开 light[k] 这个灯。
灯的颜色要想 变成蓝色 就必须同时满足下面两个条件：
1: 灯处于打开状态。
2: 排在它之前（左侧）的所有灯也都处于打开状态。

请返回能够让 所有开着的 灯都 变成蓝色 的时刻 数目 。

* n == light.length
* 1 <= n <= 5 * 10^4
* light 是 [1, 2, ..., n] 的一个排列。

### 实例
#### 1
输入：light = [2,1,3,5,4]
输出：3

#### 2
输入：light = [4,1,2,3]
输出：1