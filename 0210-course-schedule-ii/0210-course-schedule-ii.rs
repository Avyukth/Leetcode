use std::collections::{HashMap, VecDeque};

impl Solution {
    pub fn find_order(num_courses: i32, prerequisites: Vec<Vec<i32>>) -> Vec<i32> {
        
        let n = num_courses as usize;

        let mut graph:HashMap<i32, Vec<i32>> = HashMap::new();
        let mut in_degree: Vec<i32> = vec![0;n];

        for i in 0..num_courses{
            graph.entry(i).or_insert(Vec::new());
        }

        for prereq in prerequisites.iter(){
            let course = prereq[0];
            let prerequisite =prereq[1];

            graph.entry(prerequisite)
                .and_modify(|e| e.push(course))
                .or_insert(vec![course]);

            in_degree[course as usize] +=1;

        }

        let mut queue: VecDeque<i32> = VecDeque::new();
        for course in 0..num_courses{
            if in_degree[course as usize] ==0{
                queue.push_back(course);
            }
        }

        let mut result:Vec<i32> = Vec::new();

        while let Some(current) = queue.pop_front(){
            result.push(current);

            if let Some(neigh) = graph.get(&current){
                for &next_course in neigh{
                    in_degree[next_course as usize] -=1;
                    if in_degree[next_course as usize] ==0{
                        queue.push_back(next_course);
                    }
                }
            }
        }

        if result.len()!=n{
            return Vec::new();
        }
        result
    }
}