func minDays(bloomDay []int, m int, k int) int {
    n:= len(bloomDay)
    if m*k>n{
        return -1
    }
    left, right := 1, int(math.Pow(10,9))
    result:= -1
    for left<=right{
        mid :=left+(right-left)/2
        if canMake(bloomDay, m, k, mid){
            result = mid
            right = mid-1
        }else{
            left = mid+1
        }
    }
    return result
}

func canMake(bloomDay []int, m , k, day int)bool{
    count :=0
    flowers :=0
    for _, bloom:= range bloomDay{
        if bloom <= day{
            flowers++
            if flowers == k {
                count ++
                flowers =0
                if count == m{
                    return true
                } 
            }
        }else{
            flowers = 0
        }
    }
    return false
}