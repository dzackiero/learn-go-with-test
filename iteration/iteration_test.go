package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Repeat a character with how many times character repeated", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"

		if got != want {
			t.Errorf("expected '%s' but got '%s'", want, got)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}
