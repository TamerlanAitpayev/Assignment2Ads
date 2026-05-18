package main

func dfs(g *Graph, start string) []string {
    visited := map[string]bool{}
    order := []string{}
    var visit func(string)
    visit = func(u string) {
        visited[u] = true
        order = append(order, u)
        neighbors := g.adj[u]
        for i := range neighbors {
            v := neighbors[i].to
            if !visited[v] {
                visit(v)
            }
        }
    }
    visit(start)
    return order
}

func bfs(g *Graph, start string) []string {
    visited := map[string]bool{start: true}
    queue := []string{start}
    order := []string{start}
    for len(queue) > 0 {
        u := queue[0]
        queue = queue[1:]
        for _, e := range g.adj[u] {
            if !visited[e.to] {
                visited[e.to] = true
                queue = append(queue, e.to)
                order = append(order, e.to)
            }
        }
    }
    return order
}
