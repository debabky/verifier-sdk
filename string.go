package verifier_sdk

import (
	"fmt"
	"github.com/pkg/errors"
)

type String struct {
	checkName    string
	desiredValue string
}

func NewString(checkName, desiredName string) String {
	return String{
		checkName, desiredName,
	}
}

func (s *String) Verify(paramSignal []string) error {
	if paramSignal[0] != s.desiredValue {
		return errors.New(fmt.Sprintf("invalid %s", s.checkName))
	}

	return nil
}

func (s *String) GetOffset() int {
	return 1
}

func (s *String) GetName() string {
	return s.checkName
}
