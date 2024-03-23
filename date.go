package verifier_sdk

import (
	"github.com/pkg/errors"
	"strconv"
)

type Date struct {
	checkName string

	day   int
	month int
	year  int
}

// NewDate pass -1 if you want to mark that some value can be omitted
func NewDate(checkName string, day, month, year int) (Date, error) {
	if day == -1 && month == -1 && year == -1 {
		return Date{}, errors.New("failed to create a new birthday: all the values are not significant")
	}

	return Date{
		checkName, day, month, year,
	}, nil
}

func (d *Date) Verify(paramSignal []string) error {
	day, err := strconv.Atoi(paramSignal[0])
	if err != nil {
		return errors.Wrap(err, "failed to convert day signal input to int")
	}
	if d.day != -1 && day != d.day {
		return errors.New("invalid day")
	}

	month, err := strconv.Atoi(paramSignal[1])
	if err != nil {
		return errors.Wrap(err, "failed to convert month signal input to int")
	}
	if d.month != -1 && month != d.day {
		return errors.New("invalid month")
	}

	year, err := strconv.Atoi(paramSignal[2])
	if err != nil {
		return errors.Wrap(err, "failed to convert month signal input to int")
	}
	if d.year != -1 && year != d.year {
		return errors.New("invalid year")
	}

	return nil
}

func (d *Date) GetOffset() int {
	return 3
}

func (d *Date) GetName() string {
	return d.checkName
}
