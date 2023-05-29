package bench_test

import (
	"testing"
)

func itoa(i int) (s string) {
	flag := false
	if i < 0 {
		flag = true
		i = -i
	}
	if i == 0 {
		s = "0"
	}
	for j := 0; i > 0; j++ {
		s = string('0'+(i%10)) + s
		i /= 10
	}
	if flag {
		s = "-" + s
	}
	return s
}

func toString(v int) string {
	return itoa(v)
	//return fmt.Sprintf("%d", v)
	//return strconv.Itoa(v)
}

func BenchmarkToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		toString(42)
	}
}
