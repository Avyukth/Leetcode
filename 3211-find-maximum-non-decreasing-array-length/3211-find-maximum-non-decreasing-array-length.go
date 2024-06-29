func findMaximumLength(nums []int) int {
    n := len(nums)
    stk := [][3]int64{{0, 0, 0}} // last + pre, pre, dp
    var pre int64 = 0
    var res int64 = 0
    j := 0

    for i := 0; i < n; i++ {
        pre += int64(nums[i])
        j = min(j, len(stk)-1)

        for j+1 < len(stk) && pre >= stk[j+1][0] {
            j++
        }

        _, pre_pre, pre_dp := stk[j][0], stk[j][1], stk[j][2]
        cur_dp := pre_dp + 1
        res = cur_dp

        last := pre - pre_pre

        for len(stk) > 0 && stk[len(stk)-1][0] >= last+pre {
            stk = stk[:len(stk)-1]
        }

        stk = append(stk, [3]int64{last + pre, pre, cur_dp})
    }

    return int(res)

}