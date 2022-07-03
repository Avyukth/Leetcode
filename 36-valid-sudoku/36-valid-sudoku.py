class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        s=set()
        for i in range(len(board)):
            for j in range(len(board[0])):
                num=board[i][j]
                if num !=".":
                    rr=num+"r"+str(i)
                    cc=num+"c"+str(j)
                    bb = num+"b"+str(i//3)+str(j//3)
                    if (rr in s or cc in s or bb in s):
                        return False
                    else:
                        s.add(rr)
                        s.add(cc)
                        s.add(bb)
                    print(s)
        return True
            
            
            
            