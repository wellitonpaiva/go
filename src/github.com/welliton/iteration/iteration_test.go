package iteration

import "testing"

func TestIteration(t *testing.T) {
	text := IterateText("text", 5)
	expected := "texttexttexttexttext"

	if text != expected {
		t.Errorf("Expected %q, but it was %q instead", expected, text)
	}
}

func TestNumberOfTimes(t *testing.T) {
	text := IterateText("James", 3)
	expected := "JamesJamesJames"

	if text != expected {
		t.Errorf("Expected %q, but it was %q instead", expected, text)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IterateText("a", i)
	}
}
