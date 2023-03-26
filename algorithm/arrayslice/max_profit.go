package arrayslice

import "github.com/leeprince/goinfra/algorithm/pkg"

/**
给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。

在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。

返回 你能获得的 最大 利润 。



示例 1：

输入：prices = [7,1,5,3,6,4]
输出：7
解释：在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4 。
     随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6 - 3 = 3 。
     总利润为 4 + 3 = 7 。

示例 2：

输入：prices = [1,2,3,4,5]
输出：4
解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4 。
     总利润为 4 。

示例 3：

输入：prices = [7,6,4,3,1]
输出：0
解释：在这种情况下, 交易无法获得正利润，所以不参与交易可以获得最大利润，最大利润为 0 。



提示：

    1 <= prices.length <= 3 * 104
    0 <= prices[i] <= 104

作者：力扣 (LeetCode)
链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/x2zsx1/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func MaxProfit() (resp int) {
	prices := []int{7, 1, 5, 3, 6, 4}

	/*
		理解题意重点
		1. 同一天可以买入后卖出或者卖出后买入，所以原本所有的操作状态由：买入、不操作、卖出，变成：买入、卖出。最终不影响计算得到的最终收益
		2. 想收益最高，最后一天肯定不会买入，而是卖出
	*/

	/*
		回溯算法（暴力）：
		1. 状态：
			全局最大总收益maxProfitV1Resp
			第i天是否买入isBuy
			第i天操作后总收益profit
		2. 状态转换：
			回溯函数入参：每天价格、第i天、第i天是否买入、第i天操作后历史总收益；
			回溯点：i=len(prices)，非len(prices)-1

			第i天买入操作后总收益=第i-1天操作后总收益-第i天价格prices[i]；
			第i天卖出操作后总收益=第i-1天操作后总收益+第i天价格prices[i]
		3. 初始值
			参考回溯函数定义初始值：第1天，第1天不买入，第1天操作（不买入）后总收益0
		3. 输出值
			达到回溯点时：全局总收益=max(全局总收益,最后一天操作后总收益)
	*/
	resp = maxProfitV1(prices)

	/*
		动态规划：
		1. 定义状态：
			第i天操作的两种的收益，p[i][j] i第几天；j是否买入，取最大收益作为第i天操作后总收益
			p[i][0]：第i天卖出(不买入)后总收益
			p[i][1]：第i天买入后总收益
		2. 状态转换：
			前一天的买入才能卖出；前一天的卖出才能买入

			第i天卖出后总收益：p[i][0] = max(前一天买入时总收益+第i天卖出价格=p[i-1][1]+prices[i], 前一天卖出时总收益p[i-1][0])
			第i天买入后总收益：p[i][1] = max(前一天卖出时总收益-第i天买入价格=p[i-1][0]-prices[i], 前一天买入时总收益p[i-1][1])
		3. 初始化：
			初始化第一天
			第一天操作后收益：p[0][0] = 0， 还没买入就卖出收益为0
			第一天操作后收益：p[0][1] = -prices[0]，第一天买入需要花钱的
		4. 输出值：
			最后一天卖出的总收益>=最后一天买入的总收益，所以输出 p[len(prices)][0]
	*/
	resp = maxProfitV2(prices)

	/*
		动态规划优化空间：

		> 优化前
		// 第i天卖出的最大收益: max(前一天的买入才能卖出, 前一天的卖出的最大收益)
		profit[i] = append(profit[i], pkg.Max(profit[i-1][1]+prices[i], profit[i-1][0]))
		// 第i天买入的最大收益： max(前一天的卖出才能卖出, 前一天的买入的最大收益)
		profit[i] = append(profit[i], pkg.Max(profit[i-1][0]-prices[i], profit[i-1][1]))

		> 优化后
		p0, p1 = pkg.Max(p1+prices[i], p0), pkg.Max(p0-prices[i], p1)

		注意到上面的状态转移方程中，每一天的状态只与前一天的状态有关，而与更早的状态都无关，因此我们不必存储这些无关的状态，只需要将 dp[i−1][0]\textit{dp}[i-1][0]dp[i−1][0] 和 dp[i−1][1]\textit{dp}[i-1][1]dp[i−1][1] 存放在两个变量中
	*/
	resp = maxProfitV3(prices)

	/*
		贪心算法：
		1. 定义状态：
			今天价格
			明天价格
		2. 状态转换：
			明天价格比今天价格高，今天就买入明天卖出=今天价格比昨天价格高，昨天就买入今天卖出

			明天价格比今天价格高，今天就买入明天卖出
			prices[i+1]>prices[i]
			profit+(prices[i+1]+prices[i])
		3. 初始化：

		4. 输出值
			profit
	*/
	resp = maxProfit(prices)

	return
}

var maxProfitV1Resp int

func maxProfitV1(prices []int) (resp int) {
	// 传入参数特殊情况处理
	if len(prices) <= 1 {
		return 0
	}

	maxProfitV1Resp = 0

	// 调用回溯函数
	maxProfitV1Backtracking(prices, 0, false, 0)
	return maxProfitV1Resp
}

// 回溯函数
//  @param prices 每天价格
//  @param index 第i天
//  @param isBuy 第i天是否买入
//  @param profit 第i天操作后总收益
//
func maxProfitV1Backtracking(prices []int, index int, isBuy bool, profit int) {
	// 回溯点
	if len(prices) == index {
		// 输出值
		maxProfitV1Resp = pkg.Max(maxProfitV1Resp, profit)
		return
	}

	maxProfitV1Backtracking(prices, index+1, isBuy, profit)

	// 当天都会有两种操作可能
	if !isBuy {
		maxProfitV1Backtracking(prices, index+1, true, profit-prices[index])
	} else {
		maxProfitV1Backtracking(prices, index+1, false, profit+prices[index])
	}
}

func maxProfitV2(prices []int) (resp int) {
	// 传入参数特殊情况处理，无需处理
	// if len(prices) <= 1 {
	// 	return 0
	// }

	// 初始化
	profit := make([][]int, len(prices))
	profit[0] = append(profit[0], 0)
	profit[0] = append(profit[0], -prices[0])

	for i := 1; i < len(prices); i++ {
		// 状态转移
		// 第i天卖出的最大收益: max(前一天的买入才能卖出, 前一天的卖出的最大收益)
		profit[i] = append(profit[i], pkg.Max(profit[i-1][1]+prices[i], profit[i-1][0]))
		// 第i天买入的最大收益： max(前一天的卖出才能卖出, 前一天的买入的最大收益)
		profit[i] = append(profit[i], pkg.Max(profit[i-1][0]-prices[i], profit[i-1][1]))
	}

	return profit[len(prices)-1][0]
}

func maxProfitV3(prices []int) (resp int) {
	// 传入参数特殊情况处理，无需处理
	// if len(prices) <= 1 {
	// 	return 0
	// }

	// 初始化: p0:第一天卖出的收益；p1:第一天买入收益
	p0, p1 := 0, -prices[0]

	for i := 1; i < len(prices); i++ {
		// 状态转移
		p0, p1 = pkg.Max(p1+prices[i], p0), pkg.Max(p0-prices[i], p1)
	}

	return p0
}

func maxProfit(prices []int) (resp int) {
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			resp += prices[i+1] - prices[i]
		}
	}
	return
}
