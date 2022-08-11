import heapq
class Solution:
    def kSmallestPairs(self, nums1: List[int], nums2: List[int], k: int) -> List[List[int]]:
        
        out=[]
        heapq.heapify(out)
        for i in nums1:
            for j in nums2:
                if len(out)<k:
                    heapq.heappush(out, [-(i+j), [i, j]])
                else:
                    if out[0][0] < -(i+j):
                        heapq.heappop(out)
                        heapq.heappush(out, [-(i+j), [i, j]])
                    else:
                        break
        final=[]
        while len(out):
            temp=heapq.heappop(out)
            final.append(temp[1])
        return final