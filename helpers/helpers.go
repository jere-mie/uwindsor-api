package helpers

import (
	distance "github.com/lithammer/fuzzysearch/fuzzy"
)

func Similar(a string, b string) float64 {
	return 1.0 / float64(distance.LevenshteinDistance(a, b))
}

func ClosestMatch(query string, strings []string) string {
	closestString := strings[0]
	closestRatio := Similar(closestString, query)
	for _, n := range strings {
		similarity := Similar(n, query)
		if similarity > closestRatio {
			closestRatio = similarity
			closestString = n
		}
	}
	return closestString
}
