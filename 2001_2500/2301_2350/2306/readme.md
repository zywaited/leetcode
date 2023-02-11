## [2306. 公司命名](https://leetcode.cn/problems/naming-a-company/)

### 说明
给你一个字符串数组 ideas 表示在公司命名过程中使用的名字列表。公司命名流程如下：

从 ideas 中选择 2 个 不同 名字，称为 ideaA 和 ideaB 。
交换 ideaA 和 ideaB 的首字母。
如果得到的两个新名字 都 不在 ideas 中，那么 ideaA ideaB（串联 ideaA 和 ideaB ，中间用一个空格分隔）是一个有效的公司名字。
否则，不是一个有效的名字。
返回 不同 且有效的公司名字的数目

### 提示：
* 2 <= ideas.length <= 5 * 10^4
* 1 <= ideas[i].length <= 10
* ideas[i] 由小写英文字母组成
* ideas 中的所有字符串 互不相同

### 实例
#### 1
输入：ideas = ["coffee","donuts","time","toffee"]
输出：6

#### 2
输入：ideas = ["lack","back"]
输出：0