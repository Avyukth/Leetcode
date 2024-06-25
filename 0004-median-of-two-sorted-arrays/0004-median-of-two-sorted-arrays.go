func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums1)> len(nums2){
        return findMedianSortedArrays(nums2, nums1)
    }
    x,y := len(nums1), len(nums2)
    low, high := 0,x
    for low <= high{
        partX:= (low+high)/2
        partY:= (x+y+1)/2 - partX
        maxLeftX := math.Inf(-1)
        if partX !=0{
            maxLeftX = float64(nums1[partX-1])
        }
        
        minRightX := math.Inf(1)
        if partX !=x {
            minRightX = float64(nums1[partX])
        }

        maxLeftY := math.Inf(-1)
        if partY != 0 {
            maxLeftY = float64(nums2[partY-1])
        }
        minRightY := math.Inf(1)
        if partY != y{
            minRightY = float64(nums2[partY])
        }

        if  maxLeftX <= minRightY && maxLeftY <= minRightX{
            if(x+y)%2 ==0{
                return (math.Max(maxLeftX, maxLeftY) + math.Min(minRightX, minRightY)) / 2
            }else{
            return math.Max(maxLeftX, maxLeftY)
            } 
        }else if maxLeftX > minRightY {
            high = partX - 1
        } else{
            low = partX +1
        }
    }
    return 0
}