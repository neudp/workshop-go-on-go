package vanila

import (
	"os"
	"testing"
)

func BenchmarkApp_GetCharacter(b *testing.B) {
	b.ReportAllocs()

	_ = os.Setenv("MIN_LOG_LEVEL", "ERROR")

	for i := 0; i < b.N; i++ {
		app, err := NewApp()

		if err != nil {
			b.Error(err)
		}

		if _, err := app.Hadle("1"); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkApp_GetCharacterParallel(b *testing.B) {
	b.ReportAllocs()

	_ = os.Setenv("MIN_LOG_LEVEL", "ERROR")

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			app, err := NewApp()

			if err != nil {
				b.Error(err)
			}

			if _, err := app.Hadle("1"); err != nil {
				b.Error(err)
			}
		}
	})
}
