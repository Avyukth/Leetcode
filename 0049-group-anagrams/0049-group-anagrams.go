func groupAnagrams(strs []string) [][]string {
    anaMap := make(map[string][]string)

    for _, str := range strs {
        chars := strings.Split(str, "")
        sort.Strings(chars)
        sortedStr := strings.Join(chars, "")
        anaMap[sortedStr] = append(anaMap[sortedStr], str)
    }
    result  := make([][]string, 0, len(anaMap))
    for _, group := range anaMap {
        result = append(result, group)
    }
    return result 
}