package go_collections

import "math/rand"

func Reverse(list []interface{}) {
	for i := 0; i <= len(list)/2-1; i++ {
		mirror := len(list) - 1 - i
		(list)[i], (list)[mirror] = (list)[mirror], (list)[i]
	}
}

func Shuffle(list []interface{}) {
	for i := 0; i < len(list)-1; i++ {
		randomIndex := rand.Intn(i + 1)
		(list)[i], (list)[randomIndex] = (list)[randomIndex], (list)[i]
	}
}
