## [2517. 礼盒的最大甜蜜度](https://leetcode.cn/problems/maximum-tastiness-of-candy-basket/description/)
//
//### 说明
//给你一个正整数数组 price ，其中 price[i] 表示第 i 类糖果的价格，另给你一个正整数 k 。
//
//商店组合 k 类 不同 糖果打包成礼盒出售。礼盒的 甜蜜度 是礼盒中任意两种糖果 价格 绝对差的最小值。
//
//返回礼盒的 最大 甜蜜度。
//
//### 提示：
//* 1 <= price.length <= 105
//* 1 <= price[i] <= 109
//* 2 <= k <= price.length
//
//### 实例
//#### 1
//输入：price = [13,5,1,8,21,2], k = 3
//输出：8
//
//#### 2
//输入：price = [1,3,1], k = 2
//输出：2