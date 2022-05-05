
import math
import os
import random
import re
import sys

# Complete the countTriplets function below.
def countTriplets(arr, r):
    m = dict()
    a = []
    res = 0
    if r == 0 or len(arr) == 0:
        return 0
    for i, v in enumerate(arr):
        if v in m:
            m[v].append(i)
        else:
            m[v] = [i]
            a.append(v)        
    print(m)
    print(a)
    for idx, v in enumerate(arr):
        if idx == len(arr) - 2:
            break
        if r == 1 and v in m and len(m[v]) >= 3:
            l = len(m[v])
            res += int(l *(l-1) * (l-2)/ 6)
            m.pop(v)
            continue
        c1 = 0
        c2 = 0
        if v*r in m:
            _idx = 0
            po = m[v*r][_idx]
            while po < idx and _idx < len(m[v*r]) - 1:
                _idx += 1
                po = m[v*r][_idx]
            c1+= len(m[v*r]) - _idx
        if c1 == 0:
            continue
        if v*r*r in m:
            _idx = 0
            po = m[v*r*r][_idx]
            while po < idx and _idx < len(m[v*r*r]) - 1:
                _idx += 1
                po = m[v*r*r][_idx]
            c2+= len(m[v*r*r]) - _idx
        if c2 == 0:
            continue
        res += c1*c2
    return res
        
def countTriplets_2(arr, r):
    m = dict()
    res = 0
    if r == 0 or len(arr) == 0:
        return 0
    for i, v in enumerate(arr):       
        if v % r == 0:
            if v % r**2 == 0:
                l = "%d-%d" % (int(v/(r**2)), int(v/r))
                if l in m:
                    print("%d-%d-%d" % (int(v/(r**2)), int(v/r), int(v)))
                    res += m[l]
            if int(v/r) in m:
                c = m[int(v/r)]
                l = "%d-%d" % (int(v/r), v)
                if l in m:    
                    m[l] += 1 * c
                else:
                    m[l] = 1 * c
        if v in m:
            m[v] += 1
        else:
            m[v] = 1
    return res

if __name__ == '__main__':
    fptr = open("out.log", 'w')

    finput = open("input.log")

    data = finput.readlines()

    n = int(data[0].split(" ")[0])

    r = int(data[0].split(" ")[1])

    arr = list(map(int, data[1].split()))

    ans = countTriplets_2(arr, r)

    fptr.write(str(ans) + '\n')

    fptr.close()
