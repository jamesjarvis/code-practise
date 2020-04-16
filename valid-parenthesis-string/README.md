# [Valid Parenthesis String](https://leetcode.com/problems/valid-parenthesis-string/)

For some reason I was convinced of using a stack based approach to begin with.

This is not the most efficient approach, and especially difficult to do with the `*` wild character.

As it turns out, a simple count is the best solution, comparing the minimum possible value of the `(` count with the maximum possible value.
