package ci2

import "testing"

func TestAdd(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Fatalf("ci2-known-red: Add(2,3) = %d; want literal 5", got)
	}
}
