type NumArray struct {
    nums[] int
    bit []int
}


func Constructor(nums []int) NumArray {
    n:= len(nums)
    na := NumArray{
        nums: make([]int, n),
        bit : make([]int, n+1),
    }
    for i, num := range nums{
        na.Update(i, num)
    }
    return na
}


func (this *NumArray) Update(index int, val int)  {
    diff := val - this.nums[index]
    this.nums[index] = val
    index++

    for index <len(this.bit){
        this.bit[index]+= diff
        index += index & (-index) 
    }
}


func (this *NumArray) SumRange(left int, right int) int {
    return this.getSum(right) - this.getSum(left-1)
}

func (this *NumArray) getSum(index int) int {
    sum:= 0
    index++
    for index>0{
        sum += this.bit[index]
        index -= index &(-index)
    }
    return sum 
}



/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */