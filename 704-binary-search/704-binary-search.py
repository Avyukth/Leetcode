class Solution:
    def search(self, nums: List[int], target: int) -> int:
        
        lt=len(nums)
        
        l=0
        r=lt-1
        
        while l<=r:
            mid=l+(r-l)//2
            if nums[mid]==target:
                return mid
            elif nums[mid]<=target:
                l=mid+1
            else:
                r=mid-1
            
        return -1