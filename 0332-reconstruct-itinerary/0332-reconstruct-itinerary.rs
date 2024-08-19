use std::collections::{HashMap, BinaryHeap};
use std::cmp::Reverse;

type Graph = HashMap<String, BinaryHeap<Reverse<String>>>;

impl Solution {
    pub fn find_itinerary(tickets: Vec<Vec<String>>) -> Vec<String> {
        let mut graph = Self::build_graph(tickets);
        let mut route = Vec::new();

        Self::dfs("JFK", &mut graph, &mut route);
        route.reverse();
        route
    }
    fn build_graph(tickets : Vec<Vec<String>>) -> Graph{
        let mut graph = HashMap::new();
        for ticket in tickets{
            graph.entry(ticket[0].clone())
                .or_insert(BinaryHeap::new())
                .push(Reverse(ticket[1].clone()));
        }
        graph
    }
    fn dfs(airport : &str , graph : &mut Graph, route: &mut Vec<String>){
        while let Some(dest) = graph.get_mut(airport).and_then(|x| x.pop()){
            Self::dfs(&dest.0,  graph, route);
        }
        route.push(airport.to_string());
    }
}