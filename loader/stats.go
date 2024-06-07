package loader

import (
	"math"
)

// RunningStats provides Welford's algorithm for calculating running mean and variance
type RunningStats struct {
	n    int64   // number of samples
	mean float64 // running mean
	m2   float64 // squared distance from the mean
}

func NewRunningStats() *RunningStats {
	return &RunningStats{}
}

func (rs *RunningStats) Update(x float64) {
	rs.n++
	delta := x - rs.mean
	rs.mean += delta / float64(rs.n)
	delta2 := x - rs.mean
	rs.m2 += delta * delta2
}

func (rs *RunningStats) Variance() float64 { // returns variance in nanoseconds
	if rs.n < 2 {
		return 0
	}
	return rs.m2 / float64(rs.n-1)
}

func (rs *RunningStats) StdDev() float64 {
	return math.Sqrt(rs.Variance())
}

func (rs *RunningStats) Mean() float64 {
	return rs.mean
}

// UpdateFrom allows aggregating the stats from other RunningStats instances
func (rs *RunningStats) UpdateFrom(other *RunningStats) {
	if other.n == 0 {
		return
	}

	newN := rs.n + other.n
	delta := other.mean - rs.mean
	newTotal := ((rs.mean * float64(rs.n)) + (other.mean * float64(other.n))) / float64(newN)
	newM2 := rs.m2 + other.m2 + (delta * delta * float64(rs.n) * float64(other.n) / float64(newN))

	rs.m2 = newM2
	rs.mean = newTotal
	rs.n = newN
}
