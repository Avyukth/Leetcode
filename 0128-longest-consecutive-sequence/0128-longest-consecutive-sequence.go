func longestConsecutive(nums []int) int {
    numSet := make(map[int]bool)
    for _ , num := range nums{
        numSet[num] = true
    }
    longS:= 0
    for num := range numSet{
        if !numSet[num-1]{
            currNum := num 
            currS := 1
            for numSet[currNum+1]{
                currNum ++
                currS ++
            } 
            if currS>longS{
                longS = currS
            }

        }
    }
    return longS
}