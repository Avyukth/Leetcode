func maxRepeating(sequence string, word string) int {
    if !strings.Contains(sequence, word){
        return 0 
    }
    k:=1
    for strings.Contains(sequence, strings.Repeat(word, k)){
        k++
    }
    return k -1
}