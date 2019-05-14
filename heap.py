import math

class Heap:

    def __init__(self, max_length, max_heap=True):
        self.__arr = [-1] * max_length
        self.__heap_size = 0
        self.__is_max_heap = max_heap

    def size(self):
        return self.__heap_size
    
    def compare(self, val1, val2):
        if self.__is_max_heap:
            return val1 >= val2
        else:
            return val1 < val2

    def add_to_heap(self, val):
        '''
        - Add to the end
        - Sift up
        '''
        cur_position = self.__heap_size
        self.__arr[cur_position] = val
        self.__heap_size += 1
        parent_position = math.floor((cur_position-1)/2)
        while parent_position > -1 and self.compare(self.__arr[cur_position], self.__arr[parent_position]):
            tmp = self.__arr[parent_position]            
            self.__arr[parent_position] = self.__arr[cur_position]
            self.__arr[cur_position] = tmp
            cur_position = parent_position
            parent_position = math.floor((cur_position-1)/2)

    def heapify(self, idx):
        '''
        Restore heap property at index :idx
        '''
        cur_node = idx
        while self.left_index(cur_node) < self.__heap_size:
            largest_index = cur_node
            if self.compare(self.__arr[self.left_index(cur_node)], self.__arr[largest_index]):
                largest_index = self.left_index(cur_node)
            if self.right_index(cur_node) < self.__heap_size:
                if self.compare(self.__arr[self.right_index(cur_node)], self.__arr[largest_index]):
                    largest_index = self.right_index(cur_node)
            if largest_index == cur_node:
                break
            else:
                # swap down
                tmp = self.__arr[cur_node]
                self.__arr[cur_node] = self.__arr[largest_index]
                self.__arr[largest_index] = tmp
                cur_node = largest_index

    def build_heap_from_arr(self, binary_tree_arr):
        self.__heap_size = len(binary_tree_arr)
        self.__arr = binary_tree_arr
        for idx in range(math.floor(self.__heap_size/2), -1, -1):
            self.heapify(idx)

    def extract_root(self):
        # swap with the end element
        self.__heap_size -= 1
        tmp = self.__arr[self.__heap_size]
        self.__arr[self.__heap_size] = self.__arr[0]
        self.__arr[0] = tmp
        # heapify
        self.heapify(0)
        return self.__arr[self.__heap_size]

    def get_root(self):
        return self.__arr[0]

    def left_index(self, cur_node):
        return cur_node * 2 + 1
    
    def right_index(self, cur_node):
        return cur_node * 2 + 2
    
    def __repr__(self):
        return str(self.__arr[:self.__heap_size])