package topwords

import (
	"strings"
)

// return index of min element
func findMinIndex(frequencies []int) int {
	minInd := 0
	minVal := frequencies[0]
	for ind, val := range frequencies {
		if minVal > val {
			minVal = val
			minInd = ind
		}
	}
	return minInd
}

func TopWords(s string, n int) []string {

	words := strings.Split(s, " ")
	dic := make(map[string]int, len(words))
	for _, word := range words {
		dic[word]++
	}

    // get most n frequencies
	frequencies := make([]int, 0, n)
	count := 0
	for _, val := range dic {
		if count <= n{
			count++
		}
		if count < n + 1{
			frequencies = append(frequencies, val)
		} else {
			minInd := findMinIndex(frequencies)
			if frequencies[minInd] < val {
				frequencies[minInd] = val
			}
		}
	}

	res := make([]string, 0, n)
	// if dic[key]==frequency -> add key to result
	for _, v := range frequencies {
		for key := range dic {
			if dic[key] == v {
				res = append(res, key)
				dic[key] = 0
				break
			}
		}
	}

	return res
}
