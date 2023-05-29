package unpack

import (
	"errors"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

// we have four types of tokens
// S - symbol
// D - digit
// E - escape symbol
// ES - special character

func isSymbol(s rune) bool {
	return !unicode.IsDigit(s) && s != '\\'
}

func isSpecialEscapeSymbol(s rune) bool {
	return !isSymbol(s)
}

// For our compressedString we have special rules:
// 1. STRING -> <S><D><STRING_TAIL>
// 2. STRING -> <S><STRING_TAIL>
// 3. STRING -> <E><ES><D><STRING_TAIL>
// 4. STRING -> <E><ES><STRING_TAIL>
// non-compliance with these rules leads to an error

func Unpack(input string) (string, error) {
	// Place your code here.
	var compressedString = []rune(input)
	var decomprString []rune = make([]rune, 0)
	var ind int = 0
	for ind < len(compressedString) {
		if isSymbol(compressedString[ind]) { // 1 and 2 rules
			hasDigit := 0
			cnt := 1
			// defining of digit token
			if ind+1 < len(compressedString) && unicode.IsDigit(compressedString[ind+1]) {
				cnt = int(compressedString[ind+1] - '0')
				hasDigit = 1
			}
			//decompose string
			for i := 0; i < cnt; i++ {
				decomprString = append(decomprString, compressedString[ind])
			}
			ind = ind + 1 + hasDigit
		} else if compressedString[ind] == '\\' { // 3 and 4 rules
			if ind+1 < len(compressedString) && isSpecialEscapeSymbol(compressedString[ind+1]) {
				hasDigit := 0
				cnt := 1
				// defining of digit token
				if ind+2 < len(compressedString) && unicode.IsDigit(compressedString[ind+2]) {
					cnt = int(compressedString[ind+2] - '0')
					hasDigit = 1
				}
				// decompose string
				for i := 0; i < cnt; i++ {
					decomprString = append(decomprString, compressedString[ind+1])
				}
				ind = ind + 1 + hasDigit
			} else if ind+1 < len(compressedString) {
				return "", ErrInvalidString
			}
			ind++
		} else {
			return "", ErrInvalidString
		}
	}
	return string(decomprString), nil
}
