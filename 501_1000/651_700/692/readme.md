## [692. 前K个高频单词](https://leetcode-cn.com/problems/top-k-frequent-words/)

### 说明
给一非空的单词列表，返回前 k 个出现次数最多的单词。

返回的答案应该按单词出现频率由高到低排序。如果不同的单词有相同出现频率，按字母顺序排序。

* 假定 k 总为有效值， 1 ≤ k ≤ 集合元素数。
* 输入的单词均由小写字母组成。
* 尝试以 O(n log k) 时间复杂度和 O(n) 空间复杂度解决。

### 实例
#### 1
输入: ["i", "love", "leetcode", "i", "love", "coding"], k = 2
输出: ["i", "love"]

#### 2
输入: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"], k = 4
输出: ["the", "is", "sunny", "day"]