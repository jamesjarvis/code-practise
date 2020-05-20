# [House Robber](https://leetcode.com/explore/featured/card/top-interview-questions-easy/97/dynamic-programming/576/)

So, this one involves dynamic programming in order to calculate the max "money you can rob".
We have a dp array whereby the element we want to eventually retrieve is dp[len(nums)]
The first two elements are 0, nums[0] (since thats all you can do with those amount of steps)
From then on, dp[i] is the max value of the previous dp value (no change), and the value 2 steps ago + the current house value

[3, 7, 2, 6, 2, 2, 10]
[0, 1, 0, 1, 0, 0, 1] ideal choice
[3, 7, 7, 13, 13, 15, 23] dp values, ie: we choose 23
