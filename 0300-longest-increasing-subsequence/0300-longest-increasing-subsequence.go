func lengthOfLIS(nums []int) int {
    if len(nums)==0{
        return 0
    }
    tails := make([]int, len(nums))
    size :=0
    for _, num :=range nums{
        i := binaryS(tails, size, num)
        tails[i] = num
        if i == size{
            size++
        }
    }
    return size
}

func binaryS(tails []int, size, num int)int{
    left, right:=0, size
    for left<right{
        mid := left + (right-left)/2
        if tails[mid]< num{
            left = mid +1
        } else{
            right = mid
        }
    }
    return left
}