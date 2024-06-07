package loader

import (
	"testing"
)

func TestStats(t *testing.T) {
	stats := NewRunningStats()
	numbers := []float64{1, 3, 2, 4, 5, 8, 6, 7}
	for _, v := range numbers {
		stats.Update(v)
	}

	want := 6.0
	got := stats.Variance()

	if want != got {
		t.Fatalf("want %f but got %f", want, got)
	}
}

func TestStatsAggregate(t *testing.T) {
	s1 := NewRunningStats()
	s2 := NewRunningStats()
	numbers := []float64{1, 3, 2, 4, 5, 8, 6, 7}
	for _, v := range numbers[:3] {
		s1.Update(v)
	}

	for _, v := range numbers[3:] {
		s2.Update(v)
	}

	agg := NewRunningStats()
	agg.UpdateFrom(s1)
	agg.UpdateFrom(s2)

	want := 6.0
	got := agg.Variance()

	if want != got {
		t.Fatalf("want %f but got %f", want, got)
	}
}
