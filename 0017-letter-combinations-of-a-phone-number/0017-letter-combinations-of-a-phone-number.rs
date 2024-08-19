use std::collections::HashMap;
impl Solution {
    pub fn letter_combinations(digits: String) -> Vec<String> {
        if digits.is_empty(){
            return vec![];
        }
        let digit_to_letters: HashMap<char, &str> = [
        ('2', "abc"),
        ('3', "def"),
        ('4', "ghi"),
        ('5', "jkl"),
        ('6', "mno"),
        ('7', "pqrs"),
        ('8', "tuv"),
        ('9', "wxyz"),
        ]
        .iter()
        .cloned()
        .collect();
        let mut result = Vec::new();
        Self::backtrack(&digits, 0, String::new(), &mut result, &digit_to_letters );
        result
    }
    fn backtrack(digits: &str , index: usize, current : String, result : &mut Vec<String>, digit_to_letters: &HashMap<char, &str> ){
        if index == digits.len(){
            result.push(current.clone());
            return;
        } 
        if let Some(letters) = digit_to_letters.get(&digits.chars().nth(index).unwrap()){
            for letter in letters.chars(){
                let mut new_current = current.clone();
                new_current.push(letter);
                Self::backtrack(&digits, index+1, new_current, result, digit_to_letters );
            }
        }
    }
}