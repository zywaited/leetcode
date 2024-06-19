## [2813. 子序列最大优雅度](https://leetcode.cn/problems/maximum-elegance-of-a-k-length-subsequence/)

### 说明
给你一个长度为 n 的二维整数数组 items 和一个整数 k 。

items[i] = [profiti, categoryi]，其中 profiti 和 categoryi 分别表示第 i 个项目的利润和类别。

现定义 items 的 子序列 的 优雅度 可以用 total_profit + distinct_categories2 计算，其中 total_profit 是子序列中所有项目的利润总和，distinct_categories 是所选子序列所含的所有类别中不同类别的数量。

你的任务是从 items 所有长度为 k 的子序列中，找出 最大优雅度 。

用整数形式表示并返回 items 中所有长度恰好为 k 的子序列的最大优雅度。

注意：数组的子序列是经由原数组删除一些元素（可能不删除）而产生的新数组，且删除不改变其余元素相对顺序。

* 1 <= items.length == n <= 10^5
* items[i].length == 2
* items[i][0] == profit[i]
* items[i][1] == category[i]
* 1 <= profit[i] <= 10^9
* 1 <= category[i] <= n 
* 1 <= k <= n

### 实例
#### 1
输入：items = [[3,2],[5,1],[10,1]], k = 2
输出：17

#### 2
输入：items = [[3,1],[3,1],[2,2],[5,3]], k = 3
输出：19