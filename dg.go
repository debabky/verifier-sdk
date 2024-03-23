package verifier_sdk

import (
	"bytes"
	"fmt"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"math/big"
	"strings"
)

type Dg struct {
	checkName string

	digest []byte
}

func NewDg(checkName string, digest []byte) (Dg, error) {
	if len(digest) == 0 {
		return Dg{}, errors.New("failed to create a new Dg: digest is empty")
	}

	return Dg{
		checkName, digest,
	}, nil
}

func (dg *Dg) Verify(paramSignal []string) error {
	ints, err := stringsToArrayBigInt([]string{paramSignal[0], paramSignal[1]})
	if err != nil {
		return errors.Wrap(err, "failed to convert strings to big integers")
	}

	hashBytes := make([]byte, 0)
	hashBytes = append(hashBytes, ints[0].Bytes()...)
	hashBytes = append(hashBytes, ints[1].Bytes()...)

	if !bytes.Equal(dg.digest, hashBytes) {
		return errors.New("encapsulated data and proof pub signals hashes are different")
	}

	return nil
}

func stringsToArrayBigInt(publicSignals []string) ([]*big.Int, error) {
	p := make([]*big.Int, 0, len(publicSignals))
	for _, s := range publicSignals {
		sb, err := stringToBigInt(s)
		if err != nil {
			return nil, err
		}
		p = append(p, sb)
	}
	return p, nil
}

func stringToBigInt(s string) (*big.Int, error) {
	base := 10
	if bytes.HasPrefix([]byte(s), []byte("0x")) {
		base = 16
		s = strings.TrimPrefix(s, "0x")
	}
	n, ok := new(big.Int).SetString(s, base)
	if !ok {
		return nil, fmt.Errorf("can not parse string to *big.Int: %s", s)
	}
	return n, nil
}

func (dg *Dg) GetOffset() int {
	return 2
}

func (dg *Dg) GetName() string {
	return dg.checkName
}
