package googleWire

import (
	"os"
	"testing"
)

const KB = 1 << 10

func BenchmarkGetCharacter(b *testing.B) {
	b.ReportAllocs()

	_ = os.Setenv("MIN_LOG_LEVEL", "ERROR")

	for i := 0; i < b.N; i++ {
		app, err := NewApp()
		if err != nil {
			b.Error(err)
		}

		if _, err = app.Handle("1"); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkGetCharacterParallel(b *testing.B) {
	b.ReportAllocs()

	_ = os.Setenv("MIN_LOG_LEVEL", "ERROR")

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			app, err := NewApp()
			if err != nil {
				b.Error(err)
			}

			if _, err = app.Handle("1"); err != nil {
				b.Error(err)
			}
		}
	})
}
