class Solution:
    def arraySign(self, nums: List[int]) -> int:
        out=1
        for i in nums:
            out*=i
        if out>0:
            return 1
        elif out==0:
            return 0
        else:
            return -1