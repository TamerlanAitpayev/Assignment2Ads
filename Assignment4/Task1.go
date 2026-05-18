package main

import "fmt"

type Edge struct {
    to string
    w  int
}

type Graph struct {
    adj   map[string][]Edge
    verts []string
}

func NewGraph() *Graph {
    return &Graph{adj: map[string][]Edge{}}
}

func (g *Graph) add_vertex(v string) {
    if _, ok := g.adj[v]; !ok {
        g.adj[v] = nil
        g.verts = append(g.verts, v)
    }
}

func (g *Graph) add_edge(u, v string, w int) {
    g.add_vertex(u)
    g.add_vertex(v)
    g.adj[u] = append(g.adj[u], Edge{to: v, w: w})
    g.adj[v] = append(g.adj[v], Edge{to: u, w: w})
}

func (g *Graph) adjacencyList() {
    for _, v := range g.verts {
        s := v + ":"
        for _, e := range g.adj[v] {
            s += " " + e.to + "(" + fmt.Sprintf("%d", e.w) + ")"
        }
        fmt.Println(s)
    }
}
