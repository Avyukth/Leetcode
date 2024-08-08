impl Solution {
    pub fn climb_stairs(n: i32) -> i32 {
        
        if n==0{
            return 0;
        }
        if n==1{
            return 1;
        }
        if n==2{

            return 2;
        }

        let (mut  fst , mut snd) = (1,2);
        for i in 3..=n{
            let mut thd = fst+snd;
            (fst, snd) = (snd, thd);
        } 
         snd

    }
}