package vanila

import (
	"os"
	"testing"
)

func BenchmarkApp_GetCharacter(b *testing.B) {
	b.ReportAllocs()
	os.Setenv("MIN_LOG_LEVEL", "ERROR")

	for i := 0; i < b.N; i++ {
		app, err := NewApp(false)

		if err != nil {
			b.Error(err)
		}

		if _, err := app.GetCharacter("1"); err != nil {
			b.Error(err)
		}
	}
}
