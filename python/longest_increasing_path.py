"""
Longest increasing path in a matrix problem (https://leetcode.com/problems/longest-increasing-path-in-a-matrix)
"""
class Solution:
    
    def longestIncreasingPath(self, matrix: List[List[int]]) -> int:
        '''
        return the length of longest increasing path and its path node
        '''
        if len(matrix) == 0: return 0
        self.memo = [[-1 ] *len(matrix[0]) for j in range(len(matrix))]
        self.path_memo = [[None] *len(matrix[0]) for j in range(len(matrix))]
        max_path = 0
        max_path_record = []
        for i in range(0, len(matrix)):
            for j in range(0, len(matrix[i])):
                m, r = self.dfs(matrix, i, j)
                if max_path < m:
                    max_path_record = r
                    max_path = m
        return max_path, "->".join([str(matrix[i][j]) for i, j in reversed(max_path_record)])
    
    def dfs(self, matrix, i, j):        
        if self.memo[i][j] != -1:
            return self.memo[i][j], self.path_memo[i][j]
        move_direction = ((-1, 0), (1, 0), (0, -1), (0, 1))
        max_path = 0
        max_path_record = []
        for direction in move_direction:
            if self.canMove(i, j, i+direction[0], j+direction[1], matrix):
                m, r = self.dfs(matrix, i+direction[0], j+direction[1])
                if max_path < m:
                    if type(r) == list:
                        max_path_record = r
                    else:
                        max_path_record.append(r)
                    max_path = m
                elif max_path == m:
                    max_path_record.append(r)
        self.memo[i][j] = max_path + 1
        if max_path_record:
            max_path_record.append((i, j)) 
        else: max_path_record = (i, j)
        self.path_memo[i][j] = max_path_record
        return self.memo[i][j], self.path_memo[i][j]
    
    def canMove(self, i, j, next_i, next_j, matrix: List[List[int]]):
        return -1 < next_i and next_i < len(matrix) and -1 < next_j and next_j < len(matrix[0]) and matrix[next_i][next_j] > matrix[i][j]
    