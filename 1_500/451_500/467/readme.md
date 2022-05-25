## [467. 环绕字符串中唯一的子字符串](https://leetcode.cn/problems/unique-substrings-in-wraparound-string/)

### 说明
把字符串 s 看作是 “abcdefghijklmnopqrstuvwxyz” 的无限环绕字符串，所以 s 看起来是这样的：

"...zabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcd...." .
现在给定另一个字符串 p 。返回 s 中 唯一 的 p 的 非空子串 的数量 。

### 提示
* 1 <= p.length <= 10^5
* p 由小写英文字母构成

### 实例
#### 1
输入: p = "a"
输出: 1

#### 2
输入: p = "cac"
输出: 2