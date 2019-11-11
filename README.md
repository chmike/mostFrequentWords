#Compare mostFrequentWords functions

We compare three different implementations of the function `mostFrequentWords`. 
Given a map of words and occurence count, the function returns the *n* most
frequent words sorted by decreasing occurence count. 

The general working principle is to fill a slice of *n* word-count pairs 
with the most frequent count.

The first implementation (`mostFrequentWords1`) uses the standard heap package. 
The slice is used as heap with the word with smallest count at the top. A 
word-count pair is inserted in the heap if its count value is bigger than the 
word with smallest count value in the heap. When all word-count pairs have been 
processed, the heap is sorted to return a sorted word-count list.

The second implementation (`mostFrequentWords2`) also use a heap, but we don’t 
use the standard heap package. We use our own implementation. 

The third implementation (`mostFrequentWords3`) maintains the slice sorted by 
decreasing count value. Every word-count pair is inserted in place or dropped if 
its count is smaller than the last word-count pair of the list. 


The code of `mostFrequentWords3` is the simplest and the most efficient.
