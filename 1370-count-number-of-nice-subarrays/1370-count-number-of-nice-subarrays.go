func numberOfSubarrays(nums []int, k int) int {

   count:= 0
   preCount := make(map[int]int)
   preCount[0]=1
   odC:=0
   for _, num := range nums{
    if num%2 !=0{
        odC+=1
    }
    if odC >=k{
        count += preCount[odC-k]
    }
    preCount[odC]++
   }
   return count 
}