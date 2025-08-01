package main

import "strings"

func removeBadWords(text string) string {
	substrings := strings.Split(text, " ")
	censoredStrings := []string{}
	for _, word := range substrings {
		if contains(BadWords, word) {
			censoredStrings = append(censoredStrings, CensorSymbol)
		} else {
			censoredStrings = append(censoredStrings, word)
		}
	}
	return strings.Join(censoredStrings, " ")
}

func contains(slice []string, str string) bool {
	for _, word := range slice {
		if strings.EqualFold(word, str) {
			return true
		}
	}
	return false
}
