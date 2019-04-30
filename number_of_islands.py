from collections import deque

"""
BFS solution for https://leetcode.com/problems/number-of-islands
"""

class Solution:

    def numIslands(self, grid: List[List[str]]) -> int:
        if not grid or not grid[0]:
            return 0
        self.R, self.C = len(grid), len(grid[0])
        queue = deque(maxlen=self.R * self.C)
        count = 0
        for i in range(self.R):
            for j in range(self.C):
                if grid[i][j] == '1':
                    self.visit(i, j, grid, queue)
                    count += 1
        return count

    def visit(self, i, j, grid, queue):
        queue.append((i, j))
        grid[i][j] = '0'
        ds = ((-1, 0), (1, 0), (0, -1), (0, 1))
        while queue:
            x, y = queue.popleft()
            for k, v in ds:
                n_x, n_y = x + k, y + v
                if -1 < n_x < self.R and -1 < n_y < self.C and grid[n_x][n_y] == '1':
                    grid[n_x][n_y] = '0'
                    queue.append((n_x, n_y))