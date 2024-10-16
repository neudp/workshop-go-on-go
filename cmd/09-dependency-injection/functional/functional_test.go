package functional

import (
	"os"
	"testing"
)

const KB = 1 << 10

func BenchmarkGetCharacter(b *testing.B) {
	b.ReportAllocs()

	os.Setenv("MIN_LOG_LEVEL", "ERROR")

	for i := 0; i < b.N; i++ {
		if _, err := GetCharacter("1"); err != nil {
			b.Error(err)
		}
	}
}
