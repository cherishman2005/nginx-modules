package main

import (
    "container/heap"
    "fmt"
)

// 最小堆
type PriorityQueue []int

func (pq PriorityQueue) Len() int {
    return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i] < pq[j]
}

func (pq PriorityQueue) Swap(i, j int) {
     pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(num interface{}) {
    *pq = append(*pq, num.(int))
}

func (pq *PriorityQueue) Pop() interface{} {
    //n := len(*pq)
    //val := (*pq)[n-1]
    //*pq = (*pq)[:n-1]
    //return val

    old := *pq
    n := len(old)
    x := old[n-1]
    *pq = old[0 : n-1]
    return x
}

func (pq PriorityQueue) Top() int {
    val := pq[0]
    return val
}

func findKthLargest(nums []int, k int) int {
    pq := make(PriorityQueue, 0, k)
    heap.Init(&pq)
    for _, v := range nums {
        //fmt.Printf("len=%d\n", pq.Len())
        if pq.Len() < k {
            heap.Push(&pq, v)
            continue
        }
        //fmt.Printf("len=%d\n", pq.Len())

        if pq.Top() < v {
            heap.Pop(&pq)
            heap.Push(&pq, v)
        }
    }
    return pq.Top()
}

func main() {
    nums := []int{3,2,1,5,6,4}
    k := 2
    res := findKthLargest(nums, k)
    fmt.Printf("res=%d\n", res)
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
