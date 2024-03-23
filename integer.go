package verifier_sdk

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
)

// Integer represent a constraint which could be represented as integer value.
type Integer struct {
	checkName string

	minValue int
	maxValue int
}

func NewAge(checkName string, minValue, maxValue int) Integer {
	return Integer{
		checkName, minValue, maxValue,
	}
}

// Verify checks that the provided signal is in the range between min and max value
func (i *Integer) Verify(paramSignals []string) error {
	age, err := strconv.Atoi(paramSignals[0])
	if err != nil {
		return errors.Wrap(err, "failed to convert signal input to int")
	}

	if age < i.minValue || age > i.maxValue {
		return errors.New(fmt.Sprintf("invalid %s", i.checkName))
	}

	return nil
}

func (i *Integer) GetLength() int {
	return 1
}

func (i *Integer) GetName() string {
	return i.checkName
}
