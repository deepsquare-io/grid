package wordlists

import "math/rand"

func GetRandomWords(length int) []string {
	words := make([]string, length)
	for i := 0; i < length; i++ {
		words[i] = English[rand.Intn(len(English))]
	}
	return words
}
