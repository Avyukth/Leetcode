func climbStairs(n int) int {
    if n==1{
        return 1
    }
    if n==2{
        return 2
    }
    fst:=1
    snd:=2
    third:=0
    for i :=3; i <=n; i++{
        third = fst+snd
        fst, snd =snd, third
    }
    return snd 
}