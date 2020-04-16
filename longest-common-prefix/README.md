# [Longest Common Prefix](https://leetcode.com/problems/longest-common-prefix/)

This finds the minimum length of the strings. O(N)

Then it loops this amountof times and as soon as one of the characters doesnt match, it returns the longest up to that point.
O(k*N) where k is the minimum length of the strings and N is the length of the array.
basically N^2