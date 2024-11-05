package uberFx

import (
	"os"
	"testing"
)

const KB = 1 << 10

func BenchmarkGetCharacter(b *testing.B) {
	b.ReportAllocs()

	_ = os.Setenv("MIN_LOG_LEVEL", "ERROR")

	for i := 0; i < b.N; i++ {
		if err := Do(func(app *App) error {
			_, err := app.Handle("1")
			return err
		}); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkGetCharacterParallel(b *testing.B) {
	b.ReportAllocs()

	_ = os.Setenv("MIN_LOG_LEVEL", "ERROR")

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if err := Do(func(app *App) error {
				_, err := app.Handle("1")
				return err
			}); err != nil {
				b.Error(err)
			}
		}
	})
}
