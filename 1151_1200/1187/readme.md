## [使数组严格递增](https://leetcode-cn.com/problems/make-array-strictly-increasing/)
### 说明

给你两个整数数组 arr1 和 arr2，返回使 arr1 严格递增所需要的最小「操作」数（可能为 0）。

每一步「操作」中，你可以分别从 arr1 和 arr2 中各选出一个索引，分别为 i 和 j，0 <= i < arr1.length 和 0 <= j < arr2.length，然后进行赋值运算 arr1[i] = arr2[j]。

如果无法让 arr1 严格递增，请返回 -1。

### 实例
#### 1
输入：arr1 = [1,5,3,6,7], arr2 = [1,3,2,4]
输出：1
解释：用 2 来替换 5，之后 arr1 = [1, 2, 3, 6, 7]。                                                                              解释