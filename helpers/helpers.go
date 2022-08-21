package helpers

import "github.com/schollz/closestmatch"

func ClosestMatch(query string, strings *[]string) string {
	cm := closestmatch.New(*strings, []int{2})
	return cm.Closest(query)
}
