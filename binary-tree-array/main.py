def solution(arr):
    breadth = 2
    floor = 1
    left = 0
    right = 0

    while floor < len(arr):
        level = arr[floor:(floor+breadth)]
        pivot = breadth / 2

        left += sum(level[:pivot])
        right += sum(level[pivot:])

        floor += breadth
        breadth *=2

    if left < right:
        return 'Right'
    elif right < left:
        return 'Left'
    else:
        return ''

import unittest

class Test(unittest.TestCase):

    def test_function(self):
        self.assertEqual(solution([3,6,2,9,-1,10]), 'Left')
        self.assertEqual(solution([1,4,100,5]), 'Right')
        self.assertEqual(solution([1,10,5,1,0,6]), '')
        self.assertEqual(solution([]), '')
        self.assertEqual(solution([1]), '')

t = Test()
t.test_function()