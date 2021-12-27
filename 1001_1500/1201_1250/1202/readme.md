## [1202. 交换字符串中的元素](https://leetcode-cn.com/problems/smallest-string-with-swaps/)

### 说明
给你一个字符串 s，以及该字符串中的一些「索引对」数组 pairs，其中 pairs[i] = [a, b] 表示字符串中的两个索引（编号从 0 开始）。

你可以 任意多次交换 在 pairs 中任意一对索引处的字符。

返回在经过若干次交换后，s 可以变成的按字典序最小的字符串。

* 1 <= s.length <= 10^5
* 0 <= pairs.length <= 10^5
* 0 <= pairs[i][0], pairs[i][1] < s.length
* s 中只含有小写英文字母

### 实例
#### 1
输入：s = "dcab", pairs = [[0,3],[1,2]]
输出："bacd"

#### 2
输入：s = "dcab", pairs = [[0,3],[1,2],[0,2]]
输出："abcd"