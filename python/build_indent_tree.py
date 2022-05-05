"""
input
"""
test = [
    {
        "id": 1,
        "name": "First location"
    },
    {
        "id": 2,
        "name": "Second location",
        "parentId": 1
    },
    {
        "id": 3,
        "name": "Third location",
        "parentId": 2
    },
    {
        "id": 4,
        "name": "Fourth location"
    },
    {
        "id": 5,
        "name": "Fifth location",
        "parentId": 2
    }
]

"""
output: 
-1(First location)
--2(Second location)
---3(Third location)
---5(Fifth location)
-4(Fourth location)
"""

##
roots = []
id_to_node = {}
for i in test:
    if i['id'] not in id_to_node:
        id_to_node[i['id']] = []

    if "parentId" in i:
        if i["parentId"] in id_to_node:
            id_to_node[i["parentId"]].append(id_to_node[i['id']])
        else:
            id_to_node[i["parentId"]] = [id_to_node[i['id']]]

    else:
        roots.append(i["id"])
        if i['id'] not in id_to_node:
            id_to_node[i['id']] = []
print(id_to_node)

# ##
# roots = []
# node_to_children = {}
# node_to_name = {}
# """
# complexity: O(n)
# """
# for i in test:
#     if "parentId" in i:
#         if i["parentId"] in node_to_children:
#             node_to_children[i["parentId"]].append(i["id"])
#         else:
#             node_to_children[i["parentId"]] = [i["id"]]
#     else:
#         roots.append(i["id"])
#     node_to_name[i['id']] = i['name']
# res = []
#
# """
# complexity: O(n)
# """
# def trace(space, node):
#     print("%s%d(%s)" % (space, node, node_to_name[node]))
#     if node in node_to_children:
#         for c in node_to_children[node]:
#             trace(space + "-", c)
# """
# complexity: O(n)
# """
# for root in roots:
#     trace("-", root)
#
