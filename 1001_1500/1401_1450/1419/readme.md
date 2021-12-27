## [1419. 数青蛙](https://leetcode-cn.com/problems/minimum-number-of-frogs-croaking/)

### 说明
给你一个字符串 croakOfFrogs，它表示不同青蛙发出的蛙鸣声（字符串 "croak" ）的组合。
由于同一时间可以有多只青蛙呱呱作响，所以 croakOfFrogs 中会混合多个 “croak” 。请你返回模拟字符串中所有蛙鸣所需不同青蛙的最少数目。

注意：要想发出蛙鸣 "croak"，青蛙必须 依序 输出 ‘c’, ’r’, ’o’, ’a’, ’k’ 这 5 个字母。如果没有输出全部五个字母，那么它就不会发出声音。

如果字符串 croakOfFrogs 不是由若干有效的 "croak" 字符混合而成，请返回 -1 。

* 1 <= croakOfFrogs.length <= 10^5
* 字符串中的字符只有 'c', 'r', 'o', 'a' 或者 'k'

### 实例
#### 1
输入：croakOfFrogs = "croakcroak"
输出：1

#### 2
输入：croakOfFrogs = "crcoakroak"
输出：2