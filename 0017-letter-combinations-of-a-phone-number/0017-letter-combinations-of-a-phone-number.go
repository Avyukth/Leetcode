var digitToLetters = map[byte]string{
    '2': "abc",
    '3': "def",
    '4': "ghi",
    '5': "jkl",
    '6': "mno",
    '7': "pqrs",
    '8': "tuv",
    '9': "wxyz",
}

func letterCombinations(digits string) []string {
    if len(digits) == 0{
        return []string{}
    }
    result :=  []string{}
    backtrack(digits , "" , &result, 0 )
    return result
}

func backtrack(digits string , current string , result *[]string, index int ){
    if index ==len(digits){
        *result = append(*result, current)
        return  
    }
    letters := digitToLetters[digits[index]]
    for i := 0 ; i<len(letters); i++{
        backtrack(digits , current+string(letters[i]) , result, index+1 )
    }
}