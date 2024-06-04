package iteration

import (
    "testing"
)

func TestRepeat(t *testing.T) {
    got := Repeat("c", 5)
    want := "ccccc"

    if got != want {
        t.Errorf("expected %q but got %q", want, got)
    }

    got = Repeat("c", 9)
    want = "ccccccccc"

    if got != want {
        t.Errorf("expected %q but got %q", want, got)
    }
}

func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a", 10)
    }
}
