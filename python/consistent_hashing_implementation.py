"""
Future:
- Use AVL tree for ring representation
- Use MD5 hashing
"""


class ConsistentHash:

    def __init__(self, max_key_space_length, number_of_partition, nodes):
        self.number_of_partition = number_of_partition
        self.hash_ring = {}
        self.max_key_space_length = max_key_space_length
        for node_id in nodes:
            self.add_node_to_hash_ring(node_id)

    def hash(self, key):
        return key % self.max_key_space_length

    def add_node_to_hash_ring(self, node_id):
        hashed_value = self.hash(node_id)
        if node_id not in self.hash_ring.keys():
            self.hash_ring[hashed_value] = node_id

    def remove_node_to_hash_ring(self, node_id):
        hashed_value = self.hash(node_id)
        if node_id in self.hash_ring:
            self.hash_ring.pop(hashed_value)

    def get_node_for_key(self, key):
        hashed_key = self.hash(key)
        # find the first node which hashed_key is greater than this hash_key
        for node_key in sorted(self.hash_ring.keys()):
            if hashed_key <= node_key:
                return self.hash_ring[node_key]