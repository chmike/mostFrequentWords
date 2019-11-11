package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type elem struct {
	word  string
	count int
}

type elemHeap []elem

func (h elemHeap) Len() int { return len(h) }
func (h elemHeap) Less(i, j int) bool {
	return h[i].count < h[j].count || (h[i].count == h[j].count && h[i].word > h[j].word)
}
func (h elemHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h elemHeap) Push(x interface{}) { /* not used */ }
func (h elemHeap) Pop() interface{}   { /* not used */ return nil }

// using the standard heap package
func mostFrequentWords1(m map[string]int, nbrWords int) []elem {
	h := elemHeap(make([]elem, nbrWords))
	for word, count := range m {
		if count > h[0].count || (count == h[0].count && word < h[0].word) {
			h[0] = elem{word: word, count: count}
			heap.Fix(h, 0)
		}
	}
	sort.Slice(h, func(i, j int) bool {
		return h[i].count > h[j].count || (h[i].count == h[j].count && h[i].word < h[j].word)
	})
	if len(m) < nbrWords {
		h = h[:len(m)]
	}
	return h
}

// using own heap implementation
func mostFrequentWords2(m map[string]int, nbrWords int) []elem {
	h := elemHeap(make([]elem, nbrWords))
	for word, count := range m {
		if count < h[0].count || (count == h[0].count && word > h[0].word) {
			continue
		}
		h[0] = elem{word: word, count: count}
		var i int
		for {
			j := i*2 + 1 // index of left child
			if j >= nbrWords {
				break
			}
			k := j + 1 // index of right child
			if k < nbrWords && (h[k].count < h[j].count || (h[k].count == h[j].count && h[k].word > h[j].word)) {
				j = k
			}
			if h[j].count < h[i].count || (h[j].count == h[i].count && h[j].word > h[i].word) {
				h[j], h[i] = h[i], h[j]
			}
			i = j
		}
	}
	sort.Slice(h, func(i, j int) bool {
		return h[j].count < h[i].count || (h[j].count == h[i].count && h[j].word > h[i].word)
	})
	if len(m) < nbrWords {
		h = h[:len(m)]
	}
	return h
}

// using insertion sort
func mostFrequentWords3(m map[string]int, nbrWords int) []elem {
	h := elemHeap(make([]elem, nbrWords))
	last := nbrWords - 1
	for word, count := range m {
		if count < h[last].count || (count == h[last].count && word > h[last].word) {
			continue
		}
		pos := sort.Search(nbrWords, func(i int) bool { return h[i].count < count || (h[i].count == count && h[i].word > word) })
		copy(h[pos+1:], h[pos:last])
		h[pos] = elem{count: count, word: word}
	}
	if len(m) < nbrWords {
		return h[:len(m)]
	}
	return h
}

func main() {
	m := map[string]int{"the": 123, "this": 29, "house": 4, "hold": 8, "then": 27, "other": 27, "aaa": 27, "bbb": 27, "zzz": 400, "micro": 13}
	nbrWords := 20

	fmt.Println("using standard heap package:")
	r := mostFrequentWords1(m, nbrWords)
	for i := range r {
		fmt.Println(r[i])
	}

	fmt.Println("using own heap implementation:")
	r = mostFrequentWords2(m, nbrWords)
	for i := range r {
		fmt.Println(r[i])
	}

	fmt.Println("using insertion sort:")
	r = mostFrequentWords3(m, nbrWords)
	for i := range r {
		fmt.Println(r[i])
	}
}
