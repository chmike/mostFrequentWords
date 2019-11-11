package main

import "testing"

var m = map[string]int{"the": 123, "this": 29, "house": 4, "hold": 8, "then": 27, "other": 27, "aaa": 27, "bbb": 27, "zzz": 400, "micro": 13}
var h []elem

const nbrWords = 6

func BenchmarkMostFrequentWords1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		h = mostFrequentWords1(m, nbrWords)
	}
}

func BenchmarkMostFrequentWords2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		h = mostFrequentWords2(m, nbrWords)
	}
}
func BenchmarkMostFrequentWords3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		h = mostFrequentWords3(m, nbrWords)
	}
}
