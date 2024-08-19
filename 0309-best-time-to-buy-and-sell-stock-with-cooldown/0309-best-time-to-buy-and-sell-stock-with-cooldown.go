func maxProfit(prices []int) int {
    buy := - prices[0]
    sell := 0
    cool := 0
    for _, num := range prices{
        prevBuy := buy
        prevSell := sell
        buy = max(buy , cool - num)
        sell = prevBuy + num
        cool = max(cool,prevSell)
    }
    return max(cool,sell)
}