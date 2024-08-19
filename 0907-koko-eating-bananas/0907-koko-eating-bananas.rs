impl Solution {
    pub fn min_eating_speed(piles: Vec<i32>, h: i32) -> i32 {
        let mut left = 1;
        let mut right = *piles.iter().max().unwrap();
        while left <right {
            let mid = left + (right- left)/2;
            if Self::can_eat_all(&piles, mid,  h){
                right = mid
            }else{
                left = mid+1;   
            }
        }
        left
    }
    fn can_eat_all(piles: &Vec<i32>, mid :i32,  h:i32)->bool{
        let hours: i32 = piles.iter()
                            .map(|&pile|(pile+mid-1)/mid)
                            .sum();

            hours<=h
    }
}