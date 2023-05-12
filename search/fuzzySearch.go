package search

import (
	"math"
	"strings"
)

func levenshteinDistance(s1, s2 string) int {
	r1, r2 := []rune(s1), []rune(s2)
	column := make([]int, len(r1)+1)
	for y := 1; y <= len(r1); y++ {
		column[y] = y
	}
	for x := 1; x <= len(r2); x++ {
		column[0] = x
		lastDiag := x - 1
		for y := 1; y <= len(r1); y++ {
			oldDiag := column[y]
			cost := 1
			if r1[y-1] == r2[x-1] {
				cost = 0
			}
			column[y] = min(column[y]+1, column[y-1]+1, lastDiag+cost)
			lastDiag = oldDiag
		}
	}
	return column[len(r1)]
}

func min(nums ...int) int {
	m := nums[0]
	for _, n := range nums[1:] {
		if n < m {
			m = n
		}
	}
	return m
}

func FuzzySearch(query string, concertInfo map[string]string) string {
	var bestMatch string
	var bestDistance = math.MaxInt32
	for concertName, concertUrl := range concertInfo {
		concertName = strings.ToLower(concertName)
		if strings.Contains(concertName, query) {
			distance := levenshteinDistance(query, concertName)
			if distance < bestDistance {
				bestDistance = distance
				bestMatch = concertUrl
			}
		}
	}
	return bestMatch
}
