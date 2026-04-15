package boilerplate

type AggregatorImplementation interface {
	Aggregate() ([]byte, error)
}
