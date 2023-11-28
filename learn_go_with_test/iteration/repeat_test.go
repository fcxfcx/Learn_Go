package iteration

import "testing"

func TestRepeat(t *testing.T) {
	want := "aaaaa"
	got := Repeat("a", 5)
	if want != got {
		t.Errorf("want %v but got %v", want, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 1)
	}
}
