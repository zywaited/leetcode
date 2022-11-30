## [895. 最大频率栈](https://leetcode.cn/problems/maximum-frequency-stack/description/)

### 说明
设计一个类似堆栈的数据结构，将元素推入堆栈，并从堆栈中弹出出现频率最高的元素。

实现 FreqStack 类:

FreqStack() 构造一个空的堆栈。
void push(int val) 将一个整数 val 压入栈顶。
int pop() 删除并返回堆栈中出现频率最高的元素。
如果出现频率最高的元素不只一个，则移除并返回最接近栈顶的元素。

### 提示：
* 0 <= val <= 10^9
* push 和 pop 的操作数不大于 2 * 10^4。
* 输入保证在调用 pop 之前堆栈中至少有一个元素。

### 实例
#### 1
输入：
["FreqStack","push","push","push","push","push","push","pop","pop","pop","pop"],
[[],[5],[7],[5],[7],[4],[5],[],[],[],[]]
输出：[null,null,null,null,null,null,null,5,7,5,4]