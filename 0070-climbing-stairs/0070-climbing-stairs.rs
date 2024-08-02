impl Solution {
    pub fn climb_stairs(n: i32) -> i32 {
        if n==1{
            return 1
        }
        if n==2{
            return 2
        }
        let mut  snd = 2;
        let mut fst = 1;
        for _ in 3..n+1{
            let thd = fst+snd;
            (fst, snd) = (snd, thd);
        }
        snd
    }
}