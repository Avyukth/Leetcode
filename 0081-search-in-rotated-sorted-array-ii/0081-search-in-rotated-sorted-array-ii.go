func search(nums []int, target int) bool {
    left, right := 0, len(nums)-1
    for left<= right{
        mid := left+(right-left)/2
        if nums[mid]== target{
            return true
        }
        if nums[left]<nums[mid] || (nums[left]== nums[mid] && nums[right]<nums[mid]){
            if target >= nums[left] && target < nums[mid]{
                right = mid-1
            }else{
                left = mid +1
            }
        }else if  nums[right]>nums[mid] || (nums[right]== nums[mid] && nums[left]>nums[mid]){
            if target>nums[mid] && target<=nums[right]{
                left = mid +1
            }else{
                right = mid -1
            }
        }else{
            left ++
            right --
        }
    }
    return false
}