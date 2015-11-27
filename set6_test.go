package dense

import "testing"

func TestSet6(t *testing.T) {
	s := NewSet6(2, 4, 5, 7, 60)
	if s.Count() != 5 {
		t.Errorf("Set %v is supposed to contain 5 elements, got %v", s, s.Count())
	}

	for i, v := range []int64{2, 4, 5, 7, 60} {
		if !s.Contains(v) {
			t.Errorf("Set %v is supposed to contain %v", s, v)
		}
		if n, ok := s.Ordinal(v); n != uint64(i) || !ok {
			t.Errorf("%dth element is not at ordinal %d (%v)", i, n, ok)
		}
	}
}
