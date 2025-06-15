## [1432. 改变一个整数能得到的最大差值](https://leetcode.cn/problems/max-difference-you-can-get-from-changing-an-integer/)

### 说明
给你一个整数 num 。你可以对它进行以下步骤共计 两次：

选择一个数字 x (0 <= x <= 9).
选择另一个数字 y (0 <= y <= 9) 。数字 y 可以等于 x 。
将 num 中所有出现 x 的数位都用 y 替换。
令两次对 num 的操作得到的结果分别为 a 和 b 。

请你返回 a 和 b 的 最大差值 。

注意，新的整数（a 或 b）必须不能 含有前导 0，并且 非 0。

* 1 <= num <= 10^8

### 实例
#### 1
输入：num = 555
输出：888

#### 2
输入：num = 9
输出：8