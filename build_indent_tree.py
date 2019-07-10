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

roots = []
node_to_children = {}
node_to_name = {}

for i in test:
    if "parentId" in i:
        if i["parentId"] in node_to_children:
            node_to_children[i["parentId"]].append(i["id"])
        else:
            node_to_children[i["parentId"]] = [i["id"]]
    else:
        roots.append(i["id"])
    node_to_name[i['id']] = i['name']
res = []


def trace(space, node):
    print("%s%d(%s)" % (space, node, node_to_name[node]))
    if node in node_to_children:
        for c in node_to_children[node]:
            trace(space + "-", c)

for root in roots:
    trace("-", root)
