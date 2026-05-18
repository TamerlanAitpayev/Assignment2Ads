package main

import (
    "container/heap"
    "fmt"
    "math"
)

type Item struct {
    node string
    dist int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
    *pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() any {
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[0 : n-1]
    return item
}

func dijkstra(g *Graph, start string) (map[string]int, map[string]string) {
    dist := map[string]int{}
    prev := map[string]string{}
    for _, v := range g.verts {
        dist[v] = math.MaxInt
    }
    dist[start] = 0
    pq := &PriorityQueue{}
    heap.Init(pq)
    heap.Push(pq, &Item{node: start, dist: 0})
    for pq.Len() > 0 {
        item := heap.Pop(pq).(*Item)
        if item.dist > dist[item.node] {
            continue
        }
        for _, e := range g.adj[item.node] {
            nd := item.dist + e.w
            if nd < dist[e.to] {
                dist[e.to] = nd
                prev[e.to] = item.node
                heap.Push(pq, &Item{node: e.to, dist: nd})
            }
        }
    }
    return dist, prev
}

func path(prev map[string]string, target string) []string {
    result := []string{}
    for u := target; u != ""; u = prev[u] {
        result = append([]string{u}, result...)
    }
    return result
}

func main() {
    g := NewGraph()
    g.add_edge("B", "A", 10)
    g.add_edge("C", "A", 13)
    g.add_edge("D", "A", 6)
    g.add_edge("E", "A", 3)
    g.add_edge("F", "B", 13)
    g.add_edge("F", "E", 14)
    g.add_edge("G", "E", 12)
    g.add_edge("G", "D", 8)
    g.adjacencyList()
    fmt.Println(dfs(g, "F"))
    fmt.Println(bfs(g, "F"))
    dist, prev := dijkstra(g, "D")
    for _, v := range g.verts {
        fmt.Printf("%s %d %v\n", v, dist[v], path(prev, v))
    }
}
