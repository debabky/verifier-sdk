package verifier_sdk

type Constraint interface {
	Verify(paramSignal []string) error

	// GetOffset returns the length of the signal
	GetLength() int
	GetName() string
}
