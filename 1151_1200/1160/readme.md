## [1160. 拼写单词](https://leetcode-cn.com/problems/find-words-that-can-be-formed-by-characters/)

### 说明
给你一份『词汇表』（字符串数组） words 和一张『字母表』（字符串） chars。

假如你可以用 chars 中的『字母』（字符）拼写出 words 中的某个『单词』（字符串），那么我们就认为你掌握了这个单词。

注意：每次拼写时，chars 中的每个字母都只能用一次。

返回词汇表 words 中你掌握的所有单词的 长度之和。

* 1 <= words.length <= 1000
* 1 <= words[i].length, chars.length <= 100
* 所有字符串中都仅包含小写英文字母

### 实例
#### 1
输入：words = ["cat","bt","hat","tree"], chars = "atach"
输出：6

#### 2
输入：words = ["hello","world","leetcode"], chars = "welldonehoneyr"
输出：10