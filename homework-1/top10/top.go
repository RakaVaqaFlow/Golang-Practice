package top10

import (
	"sort"
	"unicode"
)

func isWordCharacter(s rune) bool {
	return !unicode.IsSpace(s)
}

func isWordCharacterAsterisk(s rune) bool {
	return unicode.IsLetter(s) || s == '-'
}

func Top10(input string, asterisk bool) []string {

	// chose a function to determine the character of the word depending on the asterisk
	var IsWordCharacter func(rune) bool = isWordCharacter
	if asterisk {
		IsWordCharacter = isWordCharacterAsterisk
	}

	// divide text into words and count their number
	var word []rune = make([]rune, 0)
	var words map[string]int = make(map[string]int)
	for _, val := range input {
		if IsWordCharacter(val) {
			if asterisk {
				val = unicode.ToLower(val)
			}
			word = append(word, val)
		} else {
			if len(word) > 0 {
				words[string(word)]++
			}
			word = nil
		}
	}
	if len(word) > 0 {
		words[string(word)]++
	}

	// remove "-" word for special task
	if _, ok := words["-"]; ok && asterisk {
		words["-"] = 0
	}

	// extract top 10 words from map
	keys := make([]string, 0, len(words))
	for k := range words {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return words[keys[i]] > words[keys[j]] ||
			(words[keys[i]] == words[keys[j]] && keys[i] < keys[j])
	})

	if len(keys) >= 10 {
		return keys[:10]
	}
	return keys

}
