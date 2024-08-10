impl Solution {
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        let n = prices.len();
        if n<=1{
            return 0;
        }
        let mut bought = -prices[0];
        let mut sold = 0;
        let mut cooldown = 0;
        for i in 1..n{
            let pre_bought = bought;
            let prev_sold = sold;
            bought = bought.max(cooldown - prices[i]);
            sold = pre_bought + prices[i];
            cooldown = cooldown.max(prev_sold) 
        }
        sold.max(cooldown)
    }
}