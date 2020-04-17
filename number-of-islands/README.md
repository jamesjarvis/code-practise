# [Number of Islands](https://leetcode.com/explore/featured/card/30-day-leetcoding-challenge/530/week-3/3302/)

Admittedly, I had to look at https://www.geeksforgeeks.org/find-number-of-islands/?ref=lbp for inspo, but it's a great idea using a graph based approach.

```
1) Initialize all vertices as not visited.
2) Do following for every vertex 'v'.
       (a) If 'v' is not visited before, call DFSUtil(v)
       (b) Do something?

DFSUtil(v)
1) Mark 'v' as visited.
2) Do something with 'v'
3) Do following for every adjacent 'u' of 'v'.
     If 'u' is not visited, then recursively call DFSUtil(u)
```
