#!/usr/bin/python
import sys

def dependency_node_traversal(edges=[]):
    '''
    :param edges: a list of edge, eg. [[4, 6], [3, 4], [2, 3]]
    '''
    if len(edges) == 0:
        print("Your input graph is empty")
        return
    result = []
    edge_map = {}
    start_nodes = []
    # find start nodes
    for edge in edges:
        if edge[0] not in start_nodes:
            # add to start_nodes
            start_nodes.append(edge[0])
    dependent_edge_map = {}
    # convert edge to map
    for edge in edges:
        if edge[1] in start_nodes:
            # remove from start_nodes and add to edge_map
            start_nodes.remove(edge[1])
        if edge[0] not in edge_map:
            edge_map[edge[0]] = [edge[1]]
        else:            
            edge_map[edge[0]].append(edge[1])
        # add to reverse edge map
        if edge[1] not in dependent_edge_map:
            dependent_edge_map[edge[1]] = [edge[0]]
        else:
            dependent_edge_map[edge[1]].append(edge[0])

    next_node = []
    cur_node = []
    # loop through all start nodes
    while len(start_nodes) > 0:
        node = start_nodes.pop()
        cur_node.append(node)
        # if this node has dependency node
        if node in edge_map:
            depend_nodes = edge_map.pop(node)
            while len(depend_nodes) > 0:        
                depend_node = depend_nodes.pop()                    
                dependent_edge_map[depend_node].remove(node)
                # if this dependent node has no other node to depend, move it to next start node
                if len(dependent_edge_map[depend_node]) == 0:
                    next_node.append(depend_node)

        if len(start_nodes) == 0:
            start_nodes = next_node
            next_node = []
            result.append(cur_node)
            cur_node = []
        
    if len(edge_map) > 0:
        print("graph has cycle dependency!")
    else:
        print("Traversal order (left -> right) is: " + str(result))

# Run by: python dependency_node_traversal.py "[[4, 6], [3, 4], [2, 3], [3, 5], [1, 2], [1, 3], [5, 6], [7, 2]]"
# Your argument is a list of edge V, at each edge, the second element is depent to the first one, 
# or another words, the first has direct path to the second one.
if __name__ == '__main__':
    edges = eval(sys.argv[1])
    print("Your edge list: " + str(edges))
    dependency_node_traversal(edges)