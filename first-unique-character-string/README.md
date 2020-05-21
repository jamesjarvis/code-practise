# [First Unique Character in a String](https://leetcode.com/explore/featured/card/top-interview-questions-easy/127/strings/881/)

For this one, it's a simple idea, store a frequency map of each character and then loop again and find the first one with a frequency of 1.

However, since we are only working with a-z, we can swap the map type for a simple array of length 26, which is around twice as fast to retrieve elements from. Same O(2n) time complexity, but each n takes less time...
