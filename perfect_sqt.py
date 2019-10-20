import math

class Solution(object):
    
    def d(self, k, n):
        print("%d %d" % (k, n))
        if k**2 == n:
            return 1
        if k == 1:
            return 1
        if n == 0:
            return 0
        if k**2 > n:
            return self.d(k-1, n)
        else:
            return 1 + self.d(k, n-k**2)
    
    def numSquares(self, n):
        """
        :type n: int
        :rtype: int
        """
        return self.d(math.sqrt(n), n)
        