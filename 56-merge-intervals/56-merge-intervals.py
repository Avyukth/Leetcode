class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        intervals.sort()
        
        out=[]
        
        for i in intervals :
            print(i)
            if len(out)==0:
                out.append(i)
            if i[0]<= out[-1][-1] :
                out[-1][-1] =max(out[-1][-1],i[-1])
            else:
                out.append(i)
        print(out)
        return out