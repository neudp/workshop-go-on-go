package functional

import (
	"os"
	"testing"
)

func BenchmarkGetCharacter(b *testing.B) {
	b.ReportAllocs()

	_ = os.Setenv("MIN_LOG_LEVEL", "ERROR")

	for i := 0; i < b.N; i++ {
		if _, err := GetCharacter("1"); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkGetCharacterParallel(b *testing.B) {
	b.ReportAllocs()

	_ = os.Setenv("MIN_LOG_LEVEL", "ERROR")

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if _, err := GetCharacter("1"); err != nil {
				b.Error(err)
			}
		}
	})
}
