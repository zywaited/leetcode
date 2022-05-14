## [691. 贴纸拼词](https://leetcode.cn/problems/stickers-to-spell-word/)

### 说明
我们有 n 种不同的贴纸。每个贴纸上都有一个小写的英文单词。

您想要拼写出给定的字符串 target ，方法是从收集的贴纸中切割单个字母并重新排列它们。如果你愿意，你可以多次使用每个贴纸，每个贴纸的数量是无限的。

返回你需要拼出 target 的最小贴纸数量。如果任务不可能，则返回 -1 。

注意：在所有的测试用例中，所有的单词都是从 1000 个最常见的美国英语单词中随机选择的，并且 target 被选择为两个随机单词的连接。

### 提示
* n == stickers.length
* 1 <= n <= 50
* 1 <= stickers[i].length <= 10
* 1 <= target <= 15
* stickers[i] 和 target 由小写英文单词组成

### 实例
#### 1
输入：stickers = ["with","example","science"], target = "thehat"
输出：3

#### 2
输入：stickers = ["notice","possible"], target = "basicbasic"
输出：-1