package verifier_sdk

import (
	"fmt"
	"github.com/iden3/go-rapidsnark/types"
	"github.com/iden3/go-rapidsnark/verifier"
	"github.com/pkg/errors"
)

type ZkVerificationBuilder struct {
	constraints    map[string]Constraint
	constraintsLen int
}

// AddConstraints adds constraints to be verified by. Returns an error if some
// constraint already added.
func (builder *ZkVerificationBuilder) AddConstraints(constraints ...Constraint) error {
	for _, constraint := range constraints {
		if builder.constraints[constraint.GetName()] != nil {
			return errors.New(fmt.Sprintf("parameter '%s' is already in the list", constraint.GetName()))
		}

		builder.constraints[constraint.GetName()] = constraint
	}

	return nil
}

// Verify verifies given signals by provided earlier constraints by AddConstraints method.
// Also verifies provided proof with Groth16 algorithm on bn256 curve.
func (builder *ZkVerificationBuilder) Verify(piA []string, piB [][]string, piC []string, pubSignals []string, verificationKey []byte) error {
	if len(builder.constraints) == 0 {
		return errors.New("you need to specify constraints first")
	}

	var signalIndex int
	for _, constraint := range builder.constraints {
		if signalIndex+constraint.GetLength() >= len(pubSignals) {
			return errors.New("signals length is too short")
		}

		err := constraint.Verify(pubSignals[signalIndex : signalIndex+constraint.GetLength()]) // check this
		if err != nil {
			return errors.Wrap(err, "verification failed")
		}

		signalIndex += constraint.GetLength()
	}

	return verifier.VerifyGroth16(types.ZKProof{
		Proof: &types.ProofData{
			A:        piA,
			B:        piB,
			C:        piC,
			Protocol: "groth16", // Unused field, may be left empty
		},
		PubSignals: pubSignals,
	}, verificationKey)
}
