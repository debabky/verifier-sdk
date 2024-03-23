package verifier_sdk

import (
	"fmt"
	"github.com/pkg/errors"
)

// String represent a constraint which could be represented as string value.
type String struct {
	checkName    string
	desiredValue string
}

func NewString(checkName, desiredName string) String {
	return String{
		checkName, desiredName,
	}
}

// Verify checks that the provided signal corresponds desired value
func (s *String) Verify(paramSignals []string) error {
	if paramSignals[0] != s.desiredValue {
		return errors.New(fmt.Sprintf("invalid %s", s.checkName))
	}

	return nil
}

func (s *String) GetLength() int {
	return 1
}

func (s *String) GetName() string {
	return s.checkName
}
