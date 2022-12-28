package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat a five times", func(t *testing.T) {
		result := Repeat("a")
		expected := "aaaaa"

		if expected != result {
			t.Errorf("expected %q but got %q", expected, result)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
