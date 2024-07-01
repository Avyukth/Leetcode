type BrowserHistory struct {
    history []string
    current int
}


func Constructor(homepage string) BrowserHistory {
    return BrowserHistory{
        history: []string{homepage},
        current: 0,
    }
}


func (this *BrowserHistory) Visit(url string)  {
    this.history = this.history[:this.current +1]
    this.history = append(this.history, url)
    this.current ++
}


func (this *BrowserHistory) Back(steps int) string {
    this.current =max(0, this.current - steps)
    return this.history[this.current]
}


func (this *BrowserHistory) Forward(steps int) string {
    this.current = min(len(this.history)-1, this.current+steps )
    return this.history[this.current]
}

func max(a, b int)int{
    if a>b {
        return a
    }
    return b
}

func min(a, b int) int{
    if a>b {
        return b
    }
    return a
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */