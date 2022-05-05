n_city = 0
n_road = 0
p1 = 0
p2 = 0
road_map = {}
with open("data.txt") as f:
    ls = f.readlines()
    #
    tmp = ls[0].strip()
    tokens = tmp.split(" ")
    n_city = int(tokens[0])
    n_road = int(tokens[1])
    #
    tmp = ls[1].strip()
    tokens = tmp.split(" ")
    p1 = tokens[0]
    p2 = tokens[1]
    for l in ls[2:]:
        tokens = l.strip().split(" ")
        if tokens[0] not in road_map:
            road_map[tokens[0]] = [tokens[1]]
        else:
            road_map[tokens[0]].append(tokens[1])
min_length, max_length = 0, (n_road/2)
i = 0
p1_step, p2_step = 0, 0
p1_q = [p1]
p2_q = [p2]
flag = False

while p1_step <= max_length and len(p1_q) > 0 and len(p2_q) > 0:
    p1_child_q = []
    p2_child_q = []
    for s in p1_q:
        if s in road_map:
            p1_child_q.extend(road_map[s])
        for t in p2_q:
            if t in road_map:
                p2_child_q.extend(road_map[t])
            if s == t:
                flag = True
                print("Catch up with len: %d" % p1_step)
                break
    p1_step += 1
    p1_q = p1_child_q
    p2_q = p2_child_q

if not flag:
    print("Cannot catch-up")
