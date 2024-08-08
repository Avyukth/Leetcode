func maxProduct(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    maxTotal := nums[0]
    maxEndingHere := nums[0]
    minEndingHere := nums[0]
    
    for i := 1; i < len(nums); i++ {
        num := nums[i]
        tempMax := max(num, max(safeMultiply(maxEndingHere, num), safeMultiply(minEndingHere, num)))
        tempMin := min(num, min(safeMultiply(maxEndingHere, num), safeMultiply(minEndingHere, num)))
        
        maxEndingHere = tempMax
        minEndingHere = tempMin
        
        maxTotal = max(maxTotal, maxEndingHere)
    }
    
    return maxTotal
}


func safeMultiply(a, b int) int {
    product := int64(a) * int64(b)
    if product > math.MaxInt32 {
        return math.MaxInt32
    }
    if product < math.MinInt32 {
        return math.MinInt32
    }
    return int(product)
}