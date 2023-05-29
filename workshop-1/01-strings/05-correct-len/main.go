package main

func IsStringShorterOrEqualThen3(s string) bool {
	return len([]rune(s)) <= 3
}
