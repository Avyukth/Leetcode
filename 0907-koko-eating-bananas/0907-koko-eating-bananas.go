func minEatingSpeed(piles []int, h int) int {
    left , right :=1, max(piles)
    result := right
    
    for left<=right{
        mid := left +(right-left)/2
        if canFinish(piles, mid, h){
            result = mid 
            right = mid -1 
        }else{
            left = mid+1
        }
    }
    return result
}

func max(piles []int )int{
    maxVal:= 0
    for _, pile:= range piles{
        if pile > maxVal {
            maxVal = pile
        }
    }
    return maxVal
}
func canFinish(piles []int, k, h int) bool{
    hour:=0
    for _, pile:= range piles{
        hour += int(math.Ceil(float64(pile)/float64(k)))
    }
    return hour<=h 
}