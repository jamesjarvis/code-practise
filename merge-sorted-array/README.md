# [Merge Sorted Array](https://leetcode.com/explore/featured/card/top-interview-questions-easy/96/sorting-and-searching/587/)

Ok so this took me 3 submissions before I realised the edge cases I had forgotten about (0's in the middle of the sorted array and an empty second list)

Basically it goes through nums1 and compares if that current int is greater than the one at the start of nums2, if it is then it swaps the two values.
It then insertion sorts nums2 so that the new swapped value is in the right place (so the smallest value is always at the front of the list basically).

Not the fastest solution, has the potential of being O(m * n) time complexity, but O(1) space complexity.

One of the "ideal" solutions is in `ideal.go`
