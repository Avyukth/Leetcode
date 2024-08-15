class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        nums_pre=[1]+[val for val in nums]
        nums_suf=[val for val in nums]+[1]

        # print(nums_pre, nums_suf)
        for i in range(1,len(nums_pre)):
            nums_pre[i]= nums_pre[i-1]*nums_pre[i]
        
        # print(nums_pre)
        for i in range(len(nums_suf) - 2, -1, -1):
            nums_suf[i] *= nums_suf[i + 1]
        print(nums_suf)
        final =[]

        for i in range(1, len(nums) + 1):
            final.append(nums_pre[i - 1] * nums_suf[i])

        return final

