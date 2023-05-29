package main

import "testing"

const (
	userID    int64 = 1234567
	salt      int64 = 7654321
	sessionID       = "abc123ccc"
)

func BenchmarkHashUser(b *testing.B) {
	b.Run("old hash, parse by user id", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = HashUser(sessionID, userID, int64(0), salt, "user_id")
		}
	})
	b.Run("old hash, parse by user_age", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = HashUser(sessionID, userID, int64(0), salt, "user_age")
		}
	})
}
