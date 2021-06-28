package main

import (
    "fmt"
)

type Node struct {
    Name    string
    Current int
    Weight  int
}

func SmoothWrr(nodes []*Node) (best *Node) {
    if len(nodes) == 0 {
        return
    }
    total := 0
    for _, node := range nodes {
        if node == nil {
            continue
        }
        total += node.Weight
        node.Current += node.Weight
        if best == nil || node.Current > best.Current {
            best = node
        }
    }
    if best == nil {
        return
    }
    best.Current -= total
    return
}

func example() {
    nodes := []*Node{
        &Node{"a", 0, 5},
        &Node{"b", 0, 1},
        &Node{"c", 0, 1},
    }

    for i := 0; i < 7; i++ {
        best := SmoothWrr(nodes)
        if best != nil {
            fmt.Println(best.Name)
        }
    }
}

func example1() {
    nodes := []*Node{
        &Node{"aa", 0, 3},
        &Node{"bb", 0, 3},
        &Node{"cc", 0, 3},
    }

    for i := 0; i < 9; i++ {
        best := SmoothWrr(nodes)
        if best != nil {
            fmt.Println(best.Name)
        }
    }
}

func main() {
    example()
    example1()
}