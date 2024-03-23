package verifier_sdk

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
)

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

func (i *Integer) Verify(paramSignal []string) error {
	age, err := strconv.Atoi(paramSignal[0])
	if err != nil {
		return errors.Wrap(err, "failed to convert signal input to int")
	}

	if age < i.minValue || age > i.maxValue {
		return errors.New(fmt.Sprintf("invalid %s", i.checkName))
	}

	return nil
}

func (i *Integer) GetOffset() int {
	return 1
}

func (i *Integer) GetName() string {
	return i.checkName
}
