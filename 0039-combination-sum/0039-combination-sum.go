func combinationSum(candidates []int, target int) [][]int {
        var result [][]int
        var current []int
        sort.Ints(candidates)
        backtrack(candidates, current, &result, 0, target)
        return result
}

func backtrack(candidates []int, current []int, result *[][]int, start, target int){
    if target ==0 {
        temp := make([]int, len(current))
        copy(temp, current)
        *result = append(*result, temp)
        return 
    }
    for i := start; i <len(candidates) ; i++{
        if candidates[i]> target{
            break
        }
        current = append(current,candidates[i] )
        backtrack(candidates, current, result, i, target - candidates[i])
        current = current[:len(current)-1]
    }
}