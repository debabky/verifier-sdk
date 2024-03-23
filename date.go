package verifier_sdk

import (
	"github.com/pkg/errors"
	"strconv"
)

// Date represent a constraint which could be represented as date.
type Date struct {
	checkName string

	day   int
	month int
	year  int
}

// NewDate accepts arguments in which -1 can be passed in case of necessity of omitting some parameter.
// If all the arguments are -1, returns an error
func NewDate(checkName string, day, month, year int) (Date, error) {
	if day == -1 && month == -1 && year == -1 {
		return Date{}, errors.New("failed to create a new birthday: all the values are not significant")
	}

	return Date{
		checkName, day, month, year,
	}, nil
}

// Verify checks that the provided signals are integer values
func (d *Date) Verify(paramSignals []string) error {
	day, err := strconv.Atoi(paramSignals[0])
	if err != nil {
		return errors.Wrap(err, "failed to convert day signal input to int")
	}
	if d.day != -1 && day != d.day {
		return errors.New("invalid day")
	}

	month, err := strconv.Atoi(paramSignals[1])
	if err != nil {
		return errors.Wrap(err, "failed to convert month signal input to int")
	}
	if d.month != -1 && month != d.day {
		return errors.New("invalid month")
	}

	year, err := strconv.Atoi(paramSignals[2])
	if err != nil {
		return errors.Wrap(err, "failed to convert month signal input to int")
	}
	if d.year != -1 && year != d.year {
		return errors.New("invalid year")
	}

	return nil
}

func (d *Date) GetLength() int {
	return 3
}

func (d *Date) GetName() string {
	return d.checkName
}
