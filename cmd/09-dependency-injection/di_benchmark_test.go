package main

import (
	"github.com/spf13/cobra"
	vanillaFunc "goOnGo/internal/swapi-func/use-case/cobra/vanilla"
	googleWire "goOnGo/internal/swapi/use-case/cobra/google-wire"
	uberFx "goOnGo/internal/swapi/use-case/cobra/uber-fx"
	"goOnGo/internal/swapi/use-case/cobra/vanilla"
	"io"
	"os"
	"testing"
)

var (
	_ = os.Setenv("MIN_LOG_LEVEL", "ERROR")
	_ = os.Setenv("DISABLE_OUTPUT", "true")
)

func BenchmarkGetCharacterGoogleWire(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		cmd := new(cobra.Command)
		*cmd = *googleWire.Cmd()
		cmd.SetOut(io.Discard)
		cmd.SetArgs([]string{"1"})

		if err := cmd.Execute(); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkGetCharacterGoogleWireParallel(b *testing.B) {
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cmd := new(cobra.Command)
			*cmd = *googleWire.Cmd()
			cmd.SetOut(io.Discard)
			cmd.SetArgs([]string{"1"})

			if err := cmd.Execute(); err != nil {
				b.Error(err)
			}
		}
	})
}

func BenchmarkGetCharacterUberFx(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		cmd := new(cobra.Command)
		*cmd = *uberFx.Cmd()
		cmd.SetOut(io.Discard)
		cmd.SetArgs([]string{"1"})

		if err := cmd.Execute(); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkGetCharacterUberFxParallel(b *testing.B) {
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cmd := new(cobra.Command)
			*cmd = *uberFx.Cmd()
			cmd.SetOut(io.Discard)
			cmd.SetArgs([]string{"1"})

			if err := cmd.Execute(); err != nil {
				b.Error(err)
			}
		}
	})
}

func BenchmarkGetCharacterVanilla(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		cmd := new(cobra.Command)
		*cmd = *vanilla.Cmd()
		cmd.SetOut(io.Discard)
		cmd.SetArgs([]string{"1"})

		if err := cmd.Execute(); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkGetCharacterVanillaParallel(b *testing.B) {
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cmd := new(cobra.Command)
			*cmd = *vanilla.Cmd()
			cmd.SetOut(io.Discard)
			cmd.SetArgs([]string{"1"})

			if err := cmd.Execute(); err != nil {
				b.Error(err)
			}
		}
	})
}

func BenchmarkGetCharacterVanillaFunc(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		cmd := new(cobra.Command)
		*cmd = *vanillaFunc.Cmd()
		cmd.SetOut(io.Discard)
		cmd.SetArgs([]string{"1"})

		if err := cmd.Execute(); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkGetCharacterVanillaFuncParallel(b *testing.B) {
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cmd := new(cobra.Command)
			*cmd = *vanillaFunc.Cmd()
			cmd.SetOut(io.Discard)
			cmd.SetArgs([]string{"1"})

			if err := cmd.Execute(); err != nil {
				b.Error(err)
			}
		}
	})
}
