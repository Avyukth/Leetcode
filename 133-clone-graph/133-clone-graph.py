"""
# Definition for a Node.
class Node:
    def __init__(self, val = 0, neighbors = None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []
"""

class Solution:
    def cloneGraph(self, node: 'Node') -> 'Node':
        if not node :
            return node
        
        q, clone=deque([node]), {node.val:Node(node.val,[])}
        
        while q:
            curr= q.popleft()
            cur_clones=clone[curr.val]
            
            for nbr in curr.neighbors:
                if nbr.val not in clone:
                    clone[nbr.val]=Node(nbr.val,[])
                    q.append(nbr)
                cur_clones.neighbors.append(clone[nbr.val])
        
        return clone[node.val]