package main

import (
	"fmt"
	"regexp"
	"strings"
)

var vowelsSet = map[rune]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
}

func isLetter(r rune) bool {
	return !((r < 'a' || r > 'z') && (r < 'A' || r > 'Z'))
}

func isWord(s string) bool {

	for _, ch := range s {
		if !isLetter(ch) {
			return false
		}
	}
	return true
}

func getFirstVowel(s string) int {
	for i, ch := range s {
		_, ok := vowelsSet[ch]
		if ok {
			return i
		}
	}
	return 0
}

func isTitleWord(s string) bool {
	return s[0] > 'A' && s[0] < 'Z'
}

func wordBuilder(word string) string {
	vowelIndex := getFirstVowel(word)
	isTitle := isTitleWord(word)
	subWords := []string{word[vowelIndex:], word[:vowelIndex], "ay"}
	out := strings.Join(subWords, "")

	if isTitle {
		out = strings.Title(strings.ToLower(out))
	}
	return out
}

func pigLatin(s string) string {
	var b strings.Builder

	res := regexp.MustCompile("\\b").Split(s, -1)

	for _, word := range res {
		if isWord(word) {
			b.WriteString(wordBuilder(word))
		} else {
			b.WriteString(word)
		}

	}
	return b.String()
}

func main() {
	input := "This tests whether a pattern matches a string."
	fmt.Println(pigLatin(input))
	}
