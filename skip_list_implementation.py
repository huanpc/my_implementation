# Implementation of skiplist (https://en.wikipedia.org/wiki/Skip_list)


class Node:

    def __init__(self, key, level):
        self.key = key
        self.next = [None] * (level + 1)
