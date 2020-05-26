# [Find all duplicates in array](https://leetcode.com/problems/find-all-duplicates-in-an-array/)

Basically, rather than using a map or something to store the duplicate occurrence, we invert the items at the index (this only works because the max number of occurrences is 2 and the value ranges is `1 <= nums[i] <= len(nums)`)
The only additional memory used is for the returning array
