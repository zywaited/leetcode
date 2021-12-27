## [1283. 使结果不超过阈值的最小除数](https://leetcode-cn.com/problems/find-the-smallest-divisor-given-a-threshold/)

### 说明
给你一个整数数组 nums 和一个正整数 threshold  ，你需要选择一个正整数作为除数，然后将数组里每个数都除以它，并对除法结果求和。

请你找出能够使上述结果小于等于阈值 threshold 的除数中 最小 的那个。

每个数除以除数后都向上取整，比方说 7/3 = 3 ， 10/2 = 5 。

题目保证一定有解。

* 如果 h 有多有种可能的值 ，h 指数是其中最大的那个。

### 实例
#### 1
输入：nums = [1,2,5,9], threshold = 6
输出：5

#### 2
输入：nums = [2,3,5,7,11], threshold = 11
输出：3
