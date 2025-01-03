package iteration

import "testing"

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 4)
	}
}

func TestRepeating(t *testing.T) {
	given := Repeat("a", 4)
	expected := "aaaa"

	if given != expected {
		t.Errorf("Expected %q but received %q", expected, given)
	}
}
