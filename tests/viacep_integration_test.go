package tests

import (
	"fmt"
	"github.com/WalterPaes/go-cep/cep"
	"github.com/WalterPaes/go-cep/viacep"
	"testing"
)

func TestViaCepIntegration(t *testing.T) {
	t.Run("SUCESS: Get a address by cep", func(t *testing.T) {
		zipcode := cep.CEP("01001-000")
		integration := viacep.NewIntegration(zipcode)

		add, err := integration.GetAddress()
		fmt.Println(add, err)
	})
}
