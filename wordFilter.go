package main

import "strings"

func wordFilter(text string, badWords []string) string {
	words := strings.Split(text, " ")
	for i, w := range words {
		lower := strings.ToLower(w)
		for _, bad := range badWords {
			if lower == bad {
				words[i] = "****"
				break
			}
		}
	}
	return strings.Join(words, " ")
}
