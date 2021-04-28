package dist

// Erlang Distribution

// Parameters
// a = Scale parameter, a > 0
// m = Shape parameter; positive integer

// Range: 0 ≤ x ≤ ∞
// Mean: am
// Variance: a^m

type ErlangDistribution struct {
	DistributionType string
}
