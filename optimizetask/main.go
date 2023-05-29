package main

import (
	"encoding/binary"
	"fmt"

	"golang.org/x/crypto/sha3"
)

func main() {
}

// HashUser convert user (or session) ID + salt to 0-99
func HashUser(sessionID string, userID, userAge, salt int64, hashField string) int64 {
	target := fmt.Sprintf("%d", salt)

	switch hashField {
	case "user_id":
		target += fmt.Sprintf("%d", userID)
	case "user_age":
		target += fmt.Sprintf("%d", userAge)
	default:
		target += sessionID
	}
	hashStr := make([]byte, 64)
	sha3.ShakeSum256(hashStr, []byte(target))
	hash := int64(binary.BigEndian.Uint32(hashStr))

	return int64(hash % 100)
}
