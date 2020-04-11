import "sort"

func groupAnagrams(strs []string) [][]string {
    mapped := make(map[string][]string)
    for _, s := range strs {
        temp := sortString(s)
        mapped[temp] = append(mapped[temp],s)
    }
    final := [][]string{}
    for _, v := range mapped {
        final = append(final, v)
    }
    return final
}

func sortString(str string) string {
    runes := []rune(str)
    sort.Slice(runes, func(a, b int) bool {
        return runes[a] < runes[b]
    })
    return string(runes)
}