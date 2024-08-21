impl Solution {
    pub fn reverse(x: i32) -> i32 {
        let mut num = x;
        let mut rev :i32 = 0;
        while num != 0{
            let pop = num%10;
            num /=10;
            if rev > i32::MAX / 10 || (rev == i32::MAX / 10 && pop > 7) {
                return 0;
            }
            if rev < i32::MIN / 10 || (rev == i32::MIN / 10 && pop < -8) {
                return 0;
            }
            rev = rev*10+pop;

        }
        rev
    }
}