class Solution:
    def checkValid(self, matrix: List[List[int]]) -> bool:
        
        for i in range(len(matrix)):
            row=set()
            col=set()
            for j in range(len(matrix[0])):
                if (matrix[i][j] in row  or matrix[j][i] in col ):
                       return False
                row.add(matrix[i][j])
                col.add(matrix[j][i])
                       
        return True