package chapter1

import (
	"container/heap"
	"math"
)

/*
Allocates representatives to states one at a
time. See https://en.wikipedia.org/wiki/Huntington%E2%80%93Hill_method
 */
func EqualProportions(pop []int, R int) []int {
	pq := &KeyValueHeap{}
	heap.Init(pq)
	rep := []int{}

	// Give each state a representative
	for s:=0;s<len(pop);s+=1 {
		rep = append(rep, 1)
		pq.Push(KeyValue{
			key: s,
			value: float64(pop[s])/math.Sqrt(2.0),
		})
	}

	// Allocate the remaining n - R representatives
	for i:=0;i< R - len(pop);i+=1 {
		s := heap.Pop(pq)
		rep[s.(KeyValue).key] = rep[s.(KeyValue).key] + 1
		priority := float64(pop[s.(KeyValue).key]) / math.Sqrt(float64(rep[s.(KeyValue).key] * (rep[s.(KeyValue).key] + 1)))
		pq.Push(KeyValue{
			key:   s.(KeyValue).key,
			value: priority,
		})
	}

	return rep
}

type KeyValueHeap []KeyValue

type KeyValue struct {
	key int
	value float64
}

func (h KeyValueHeap) Len() int           { return len(h) }
func (h KeyValueHeap) Less(i, j int) bool { return h[i].key < h[j].key }
func (h KeyValueHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *KeyValueHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(KeyValue))
}

func (h *KeyValueHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}