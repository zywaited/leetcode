## [2166. 设计位集](https://leetcode-cn.com/problems/design-bitset/)

### 说明
位集 Bitset 是一种能以紧凑形式存储位的数据结构。

请你实现 Bitset 类。
* Bitset(int size) 用 size 个位初始化 Bitset ，所有位都是 0 。
* void fix(int idx) 将下标为 idx 的位上的值更新为 1 。如果值已经是 1 ，则不会发生任何改变。
* void unfix(int idx) 将下标为 idx 的位上的值更新为 0 。如果值已经是 0 ，则不会发生任何改变。
* void flip() 翻转 Bitset 中每一位上的值。换句话说，所有值为 0 的位将会变成 1 ，反之亦然。
* boolean all() 检查 Bitset 中 每一位 的值是否都是 1 。如果满足此条件，返回 true ；否则，返回 false 。
* boolean one() 检查 Bitset 中 是否 至少一位 的值是 1 。如果满足此条件，返回 true ；否则，返回 false 。
* int count() 返回 Bitset 中值为 1 的位的 总数 。
* String toString() 返回 Bitset 的当前组成情况。注意，在结果字符串中，第 i 个下标处的字符应该与 Bitset 中的第 i 位一致。

提示
* 1 <= size <= 10^5
* 0 <= idx <= size - 1
* 至多调用 fix、unfix、flip、all、one、count 和 toString 方法 总共 10^5 次
* 至少调用 all、one、count 或 toString 方法一次
* 至多调用 toString 方法 5 次