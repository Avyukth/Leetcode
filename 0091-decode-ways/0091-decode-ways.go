func numDecodings(s string) int {
    n:=len(s)
    if n==0. || s[0]=='0'{
        return 0
    }
    dp:= make([]int, n+1)
    dp[0]=1
    dp[1]=1
    for i :=2 ; i <=n ; i++{
        if s[i-1] !='0'{
            dp[i]=dp[i-1] 
        }
        tD , _ := strconv.Atoi(s[i-2:i])
        if tD >=10 && tD <=26 {
            
            dp[i]+=dp[i-2]
        }


    }
    return dp[n]
}