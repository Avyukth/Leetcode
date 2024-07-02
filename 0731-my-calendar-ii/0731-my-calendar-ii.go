type MyCalendarTwo struct {
     events [][2]int
     overlaps [][2]int
}



func Constructor() MyCalendarTwo {
    return MyCalendarTwo{events: make([][2]int,0),
    overlaps: make( [][2]int,0),}
}


func (this *MyCalendarTwo) Book(start int, end int) bool {
    for _, overLap :=range this.overlaps{
        if start<overLap[1] && end >overLap[0]{
            return false
        }
    }
        for _, event := range this.events{
        if start < event[1] && end > event[0]{
            this.overlaps = append(this.overlaps, [2]int{max(start, event[0]),min(end, event[1])})
        }
    } 
    this.events = append(this.events, [2]int{start, end})
    return true
}


/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */