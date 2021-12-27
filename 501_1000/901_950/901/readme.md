## [901. 股票价格跨度](https://leetcode-cn.com/problems/online-stock-span/)

### 说明
编写一个 StockSpanner 类，它收集某些股票的每日报价，并返回该股票当日价格的跨度。

今天股票价格的跨度被定义为股票价格小于或等于今天价格的最大连续日数（从今天开始往回数，包括今天）。

例如，如果未来7天股票的价格是 [100, 80, 60, 70, 60, 75, 85]，那么股票跨度将是 [1, 1, 1, 2, 1, 4, 6]。

* 调用 StockSpanner.next(int price) 时，将有 1 <= price <= 10^5。
* 每个测试用例最多可以调用  10000 次 StockSpanner.next。
* 在所有测试用例中，最多调用 150000 次 StockSpanner.next。

### 实例
#### 1
输入：["StockSpanner","next","next","next","next","next","next","next"], [[],[100],[80],[60],[70],[60],[75],[85]]
输出：[null,1,1,1,2,1,4,6]