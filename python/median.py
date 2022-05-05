from heap import Heap

testcases = [
        [[1, 3, 9, 5, 7, 5, 0, 4, 3, 2, 9, 1, 0], [1, 2, 3, 4, 5, 5, 5, 4.5, 4, 3.5, 4, 3.5, 3]],
        [[6, 1, 2], [6, 3.5, 2]],
        [[3, 9, 2, 5, 2], [3, 6, 3, 4, 3]],
        [[5, 5, 5, 5], [5, 5, 5, 5]],
        [[9, 6, 4, 3, 1], [9, 7.5, 6, 5, 4]],
        [[1, 2, 3, 4], [1, 1.5, 2, 2.5]],
        [[1, 3, 2], [1, 2, 2]],
        [[1], [1]]
]

class MedianGetter:

    def __init__(self, max_length):
        self.max_heap = Heap(max_length)
        self.min_heap = Heap(max_length, max_heap=False)
        self.median = 0

    def add(self, val):
        if val >= self.median:
            self.min_heap.add_to_heap(val)
        else: 
            self.max_heap.add_to_heap(val)
        
        if self.max_heap.size() == self.min_heap.size():
            self.median = (self.min_heap.get_root() + self.max_heap.get_root())/2

        elif self.max_heap.size() > self.min_heap.size():
            if self.max_heap.size() == self.min_heap.size() + 2:    
                new_min_root = self.max_heap.extract_root()
                self.min_heap.add_to_heap(new_min_root)
                self.median = (self.min_heap.get_root() + self.max_heap.get_root())/2
            else:
                self.median = self.max_heap.get_root()

        else:
            if self.min_heap.size() == self.max_heap.size() + 2:    
                new_max_root = self.min_heap.extract_root()
                self.max_heap.add_to_heap(new_max_root)
                self.median = (self.min_heap.get_root() + self.max_heap.get_root())/2
            else:
                self.median = self.min_heap.get_root()

    def get_median(self):
        return self.median

for test in testcases:
    median_getter = MedianGetter(50)
    res = []
    for i in test[0]:
        median_getter.add(i)
        res.append(median_getter.get_median())
    if test[1] != res:
        print("Failed testcase", test[0], test[1], res)
        
print("Done")