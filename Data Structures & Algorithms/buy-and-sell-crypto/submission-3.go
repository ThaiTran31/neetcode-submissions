func maxProfit(prices []int) int {
    maxP := 0
    minBuy := math.MaxInt32

    for _, sell := range prices {
        if sell - minBuy > maxP {
            maxP = sell - minBuy
        }
        if sell < minBuy {
            minBuy = sell
        }
    }
    return maxP
}