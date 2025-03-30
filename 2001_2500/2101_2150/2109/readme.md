## [2109. 向字符串添加空格](https://leetcode.cn/problems/adding-spaces-to-a-string/description/)

### 说明
给你一个下标从 0 开始的字符串 s ，以及一个下标从 0 开始的整数数组 spaces 。

数组 spaces 描述原字符串中需要添加空格的下标。每个空格都应该插入到给定索引处的字符值 之前 。

例如，s = "EnjoyYourCoffee" 且 spaces = [5, 9] ，那么我们需要在 'Y' 和 'C' 之前添加空格，这两个字符分别位于下标 5 和下标 9 。因此，最终得到 "Enjoy Your Coffee" 。
请你添加空格，并返回修改后的字符串。

* 1 <= s.length <= 3 * 10^5
* s 仅由大小写英文字母组成
* 1 <= spaces.length <= 3 * 10^5
* 0 <= spaces[i] <= s.length - 1
* spaces 中的所有值 严格递增

### 实例
#### 1
输入：s = "LeetcodeHelpsMeLearn", spaces = [8,13,15]
输出："Leetcode Helps Me Learn"

#### 2
输入：s = "icodeinpython", spaces = [1,5,7,9]
输出："i code in py thon"