package verifier_sdk

type Constraint interface {
	Verify(paramSignal []string) error
	GetOffset() int
	GetName() string
}
