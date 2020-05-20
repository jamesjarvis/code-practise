def solution(tree):
    length = len(tree)
    height = 0
    breadth = 1
    floor = 0

    while length > 0:
        if isEmpty(tree[floor:(floor+breadth)]):
            return height
        length -= breadth
        floor += breadth
        breadth *= 2
        height += 1
    return height
        
def isEmpty(treesnippet):
    for e in treesnippet:
        if e != -1:
            return False
    return True