## [只出现一次的数字 II](https://leetcode-cn.com/problems/single-number-ii/)
### 说明

给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。

### 实例
#### 1

输入: [2,2,3,2]
输出: 3

#### 2

输入: [0,1,0,1,0,1,99]
输出: 99

### 要求

* 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

### 实现
* one hashmap不符合要求，需要找到类似于a^a^a=0这种方法(因为二进制a^a=0, 136题([只出现一次的数字](https://github.com/zywaited/leetcode/tree/master/101-150/136/))就是这种方式)