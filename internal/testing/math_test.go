package testing

import (
	"fmt"
	"os"
	"testing"
	"unsafe"
)

/*
Для тестирования go предоставляет пакет testing
Unit-тесты должны быть в пакете с тестируемым кодом и иметь суффикс _test
Интеграционные и E2E тесты обычно находятся вне пакета internal

Тесты семантически делятся на три типа:
1. Тесты, которые проверяют корректность работы функции (unit-тесты)
   такие тесты должны иметь сигнатуру Test<Name>(t *testing.T) и возвращать void
2. Тесты, которые проверяют эффективность работы функции (benchmark-тесты)
   такие тесты должны иметь сигнатуру Benchmark<Name>(b *testing.B) и возвращать void
3. Примеры использования функции (example-тесты)
   такие тесты должны иметь сигнатуру Example<Name>() и возвращать void
*/

/*
Существует 2 основных подхода к написанию тестов на go:
1. Тест-функции
2. Сьюит-функции
*/

// Тест-функции

func TestAdd(t *testing.T) {
	t.Parallel() // помечаем тест как безопасный для параллельного выполнения

	if os.Getenv("TEST_ADD") == "SKIP" {
		t.Skip("skipping test") // пропускаем тест если переменная окружения TEST_ADD равна "SKIP"
	}

	if testing.Short() {
		t.Skip("skipping test in short mode") // пропускаем тест если запущен в коротком режиме
	}

	tests := []struct {
		name     string
		numbers  []float64
		expected float64
	}{
		{
			name:     "Add two numbers",
			numbers:  []float64{1, 2},
			expected: 3,
		},
		{
			name:     "Add three numbers",
			numbers:  []float64{1, 2, 3},
			expected: 6,
		},
		{
			name:     "Add four numbers",
			numbers:  []float64{1, 2, 3, 4},
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { // t.Run позволяет группировать тесты
			actual := Add(tt.numbers...)
			if actual != tt.expected {
				// t.Errorf позволяет записывать ошибки с форматированием
				t.Errorf("expected %v, got %v", tt.expected, actual)

				/*
					также есть другие методы:
					t.Fatalf - записывает ошибку и завершает тест
					t.Logf - записывает сообщение (не является ошибкой)

					есть еще методы-близнецы без суффикса f, которые не принимают форматирование
				*/
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		numbers  []float64
		expected float64
	}{
		{
			name:     "Subtract two numbers",
			numbers:  []float64{1, 2},
			expected: -1,
		},
		{
			name:     "Subtract three numbers",
			numbers:  []float64{1, 2, 3},
			expected: -4,
		},
		{
			name:     "Subtract four numbers",
			numbers:  []float64{1, 2, 3, 4},
			expected: -8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Subtract(tt.numbers...)
			if actual != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

//... etc

// Сьюит-функции
/*
Сьюит-функции позволяют группировать тесты по смыслу
*/
func TestMath(t *testing.T) {
	t.Parallel()

	t.Run("Add", func(t *testing.T) {
		t.Parallel()

		if os.Getenv("TEST_ADD") == "SKIP" {
			t.Skip("skipping test")
		}

		if testing.Short() {
			t.Skip("skipping test in short mode")
		}

		tests := []struct {
			name     string
			numbers  []float64
			expected float64
		}{
			{
				name:     "Add two numbers",
				numbers:  []float64{1, 2},
				expected: 3,
			},
			{
				name:     "Add three numbers",
				numbers:  []float64{1, 2, 3},
				expected: 6,
			},
			{
				name:     "Add four numbers",
				numbers:  []float64{1, 2, 3, 4},
				expected: 10,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) { // t.Run позволяет группировать тесты
				actual := Add(tt.numbers...)
				if actual != tt.expected {
					t.Errorf("expected %v, got %v", tt.expected, actual)
				}
			})
		}
	})

	t.Run("Subtract", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name     string
			numbers  []float64
			expected float64
		}{
			{
				name:     "Subtract two numbers",
				numbers:  []float64{1, 2},
				expected: -1,
			},
			{
				name:     "Subtract three numbers",
				numbers:  []float64{1, 2, 3},
				expected: -4,
			},
			{
				name:     "Subtract four numbers",
				numbers:  []float64{1, 2, 3, 4},
				expected: -8,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				actual := Subtract(tt.numbers...)
				if actual != tt.expected {
					t.Errorf("expected %v, got %v", tt.expected, actual)
				}
			})
		}
	})

	//... etc
}

// Benchmark-тесты

func benchmarkMultiply(b *testing.B, nums []float64) {
	b.ReportAllocs()                       // показывает количество аллокаций памяти
	b.SetBytes(int64(unsafe.Sizeof(nums))) // устанавливает размер данных в байтах
	b.SetParallelism(1)                    // устанавливает количество параллельных процессов
	b.ResetTimer()                         // сбрасывает таймер, чтобы измерить только время выполнения функции

	for i := 0; i < b.N; i++ {
		_ = Multiply(nums...)
	}
}

func benchmarkDivide(b *testing.B, nums []float64) {
	b.ReportAllocs()
	b.SetBytes(int64(unsafe.Sizeof(nums)))
	b.SetParallelism(1)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Divide(nums...)
	}
}

func BenchmarkMultiply10(b *testing.B) {
	nums := make([]float64, 10)
	for i := 0; i < 10; i++ {
		nums[i] = float64(i + 1)
	}

	benchmarkMultiply(b, nums)
}

func BenchmarkMultiply100(b *testing.B) {
	nums := make([]float64, 100)
	for i := 0; i < 100; i++ {
		nums[i] = float64(i + 1)
	}

	benchmarkMultiply(b, nums)
}

func BenchmarkMultiply1000(b *testing.B) {
	nums := make([]float64, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = float64(i + 1)
	}

	benchmarkMultiply(b, nums)
}

func BenchmarkDivide10(b *testing.B) {
	nums := make([]float64, 10)
	for i := 0; i < 10; i++ {
		nums[i] = float64(i + 1)
	}

	benchmarkDivide(b, nums)
}

func BenchmarkDivide100(b *testing.B) {
	nums := make([]float64, 100)
	for i := 0; i < 100; i++ {
		nums[i] = float64(i + 1)
	}

	benchmarkDivide(b, nums)
}

func BenchmarkDivide1000(b *testing.B) {
	nums := make([]float64, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = float64(i + 1)
	}

	benchmarkDivide(b, nums)
}

// Example-тесты
/*
Example-тесты позволяют показать пример использования функции
Обычно используются для документирования кода, однако они все равно выполняются.
Успешное выполнение определяется по отсутствию ошибок, а также можно использовать // Output: <expected_output>
для проверки вывода в stdout
*/

func ExampleAdd() {
	fmt.Println(Add(1, 2, 3))
	// Output: 6
}

func ExampleSubtract() {
	fmt.Println(Subtract(1, 2, 3))
	// Output: -4
}

//... etc
