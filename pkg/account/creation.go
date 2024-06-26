package account

import (
	"errors"

	"github.com/ZippyChain/zippysdk/pkg/keys"
	"github.com/ZippyChain/zippysdk/pkg/mnemonic"
	"github.com/ZippyChain/zippysdk/pkg/store"
)

var (
	AccountByNameExists = errors.New("name chosen for account already exists")
)

type Creation struct {
	Name            string
	Passphrase      string
	Mnemonic        string
	HdAccountNumber *uint32
	HdIndexNumber   *uint32
}

func New() string {
	return "New Account"
}

func IsValidPassphrase(pass string) bool {
	return true
}

// CreateNewLocalAccount assumes all the inputs are valid, legitmate
func CreateNewLocalAccount(candidate *Creation) error {
	ks := store.FromAccountName(candidate.Name)
	if candidate.Mnemonic == "" {
		candidate.Mnemonic = mnemonic.Generate()
	}
	// Hardcoded index of 0 here.
	private, _ := keys.FromMnemonicSeedAndPassphrase(candidate.Mnemonic, 0)
	_, err := ks.ImportECDSA(private.ToECDSA(), candidate.Passphrase)
	if err != nil {
		return err
	}
	return nil
}
