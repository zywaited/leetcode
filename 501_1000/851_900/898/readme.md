## [898. 子数组按位或操作](https://leetcode-cn.com/problems/bitwise-ors-of-subarrays/)

### 说明
我们有一个非负整数数组 A。

对于每个（连续的）子数组 B = [A[i], A[i+1], ..., A[j]] （ i <= j），我们对 B 中的每个元素进行按位或操作，获得结果 A[i] | A[i+1] | ... | A[j]。

返回可能结果的数量。 （多次出现的结果在最终答案中仅计算一次。）

* 1 <= A.length <= 50000
* 0 <= A[i] <= 10^9

### 实例
#### 1
输入：[0]
输出：1

#### 2
输入：[1,1,2]
输出：3