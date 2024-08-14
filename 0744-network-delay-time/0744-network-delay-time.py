class Solution:
    def networkDelayTime(self, times: List[List[int]], n: int, k: int) -> int:
        graph = [[] for _ in range(n+1)]
        for u, v, w in times :
            graph[u].append((v,w))

        dist = [float('inf')]*(n+1)
        dist[k] = 0
        pq =[(0,k)]
        while pq :
            d, node = heapq.heappop(pq)
            if d>dist[node] :
                continue
            
            for neighbour, weight in graph[node]:
                new_dist = d + weight
                if  new_dist < dist[neighbour]:
                    dist[neighbour] = new_dist
                    heapq.heappush(pq, (new_dist, neighbour))

        max_dist = max(dist[1:])
        return max_dist if max_dist < float('inf') else -1
