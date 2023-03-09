package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	exp := "aaaaa"
	if repeated != exp {
		t.Errorf("expected %q got %q", exp, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
