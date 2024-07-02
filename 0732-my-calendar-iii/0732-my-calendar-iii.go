type MyCalendarThree struct {
    events map[int]int
}


func Constructor() MyCalendarThree {
    return MyCalendarThree{ events: make(map[int]int),}
}


func (this *MyCalendarThree) Book(startTime int, endTime int) int {
    this.events[startTime]++
    this.events[endTime]--
    maxB , currB := 0,0
    times := make([]int, 0, len(this.events))
    for time := range this.events{
        times=append(times, time)
    }
    sort.Ints(times)
    for _, time:= range times{
        currB += this.events[time]
        if currB > maxB{
            maxB = currB
        }
    }
    return maxB
}


/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */