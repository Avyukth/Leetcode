impl Solution {
    pub fn search(nums: Vec<i32>, target: i32) -> i32 {
        let mut n = nums.len()as i32 -1;
        let mut start = 0;
        while start<=n{
            let mut mid = start + (n- start)/2;
            if nums[mid as usize] ==target{
                return mid as i32;
            }else if nums[mid as usize] <target{ 
                start= mid+1;
            }else{
                n= mid-1;

            }
        }
        -1
    }
}