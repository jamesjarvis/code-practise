# [Implement strStr or indexOf](https://leetcode.com/explore/featured/card/top-interview-questions-easy/127/strings/885/)

After looking at other submissions, it could have potentially been improved for legibility by using the substring snippet:

```go
for i := 0; i + len(needle) <= len(haystack); i++ {
  if haystack[i:i+len(needle)] == needle {
    return i
  }
}
```

but I think my method of just couning the matching digitst is also alright
