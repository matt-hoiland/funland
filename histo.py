#!/usr/bin/env python3

with open('data.txt') as f:
    data = f.readlines()
    data = [[int(x) for x in datum.strip().split(':')] for datum in data]
    data = [a*100 + b for a, b in data]
    histo = {
        750: [],
        800: [],
        810: [],
        820: [],
        830: [],
        840: [],
        850: [],
        900: [],
        910: [],
        920: [],
        930: [],
        940: [],
        950: [],
        1000: []
    }
    data.sort()
    print(data)
    print(min(data))
    print(max(data))

    print(list(histo.keys()))
    i = 0
    key = list(histo.keys())[i]
    for datum in data:
        if datum//10 != key//10:
            i += 1
            key = list(histo.keys())[i]
        histo[key].append(datum)

    for key in histo.keys():
        print(key, histo[key])

    print(data[len(data)//2])
    data = [datum//100 * 60 + datum % 100 for datum in data]
    average = sum(data)/len(data)
    print(average//60, average % 60)
