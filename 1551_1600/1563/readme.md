## [1563. 石子游戏 V](https://leetcode-cn.com/problems/stone-game-v/)

### 说明
几块石子 排成一行 ，每块石子都有一个关联值，关联值为整数，由数组 stoneValue 给出。

游戏中的每一轮：Alice 会将这行石子分成两个 非空行（即，左侧行和右侧行）；
Bob 负责计算每一行的值，即此行中所有石子的值的总和。
Bob 会丢弃值最大的行，Alice 的得分为剩下那行的值（每轮累加）。
如果两行的值相等，Bob 让 Alice 决定丢弃哪一行。下一轮从剩下的那一行开始。

只 剩下一块石子 时，游戏结束。Alice 的分数最初为 0 。

返回 Alice 能够获得的最大分数 。

* 1 <= stoneValue.length <= 500
* 1 <= stoneValue[i] <= 10^6

### 实例
#### 1
输入：stoneValue = [6,2,3,4,5,5]
输出：18

#### 2
输入：stoneValue = [7,7,7,7,7,7,7]
输出：28