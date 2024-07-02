type RangeFreqQuery struct {
    indic map[int][]int
}


func Constructor(arr []int) RangeFreqQuery {
    rfq := RangeFreqQuery{
        indic : make(map[int][]int),
    }
    for i, num := range arr{
        rfq.indic[num]=append(rfq.indic[num], i) 
    }
    return rfq
}


func (this *RangeFreqQuery) Query(left int, right int, value int) int {
    if _, exists:= this.indic[value]; !exists{
        return 0
    }
    indics:= this.indic[value]
    leftI := sort.SearchInts(indics, left)
    rightI := sort.SearchInts(indics, right+1)
    return rightI- leftI
}


/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */