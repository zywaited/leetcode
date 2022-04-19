## [2234. 花园的最大总美丽值](https://leetcode-cn.com/problems/maximum-total-beauty-of-the-gardens/)

### 说明
Alice 是 n 个花园的园丁，她想通过种花，最大化她所有花园的总美丽值。

给你一个下标从 0 开始大小为 n 的整数数组 flowers ，其中 flowers[i] 是第 i 个花园里已经种的花的数目。
已经种了的花 不能 移走。同时给你 newFlowers ，表示 Alice 额外可以种花的 最大数目 。
同时给你的还有整数 target ，full 和 partial 。

如果一个花园有 至少 target 朵花，那么这个花园称为 完善的 ，花园的 总美丽值 为以下分数之 和 ：

完善 花园数目乘以 full.
剩余 不完善 花园里，花的 最少数目 乘以 partial 。如果没有不完善花园，那么这一部分的值为 0 。
请你返回 Alice 种最多 newFlowers 朵花以后，能得到的 最大 总美丽值。

提示
* 1 <= flowers.length <= 10^5
* 1 <= flowers[i], target <= 10^5
* 1 <= newFlowers <= 10^10
* 1 <= full, partial <= 10^5

### 实例
#### 1
输入：flowers = [1,3,1,1], newFlowers = 7, target = 6, full = 12, partial = 1
输出：14

#### 2
输入：flowers = [2,4,5,3], newFlowers = 10, target = 5, full = 2, partial = 6
输出：30