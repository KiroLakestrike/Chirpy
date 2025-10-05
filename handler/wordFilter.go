package handler

import "strings"

func wordFilter(text string, badWords []string) string {
	words := strings.Split(text, " ")
	for i, word := range words {
		lowered := strings.ToLower(word)
		for _, bad := range badWords {
			if lowered == bad {
				words[i] = "****"
			}
		}
	}
	return strings.Join(words, " ")
}
