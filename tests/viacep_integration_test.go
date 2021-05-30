package tests

import (
	"github.com/WalterPaes/go-cep/cep"
	"github.com/WalterPaes/go-cep/viacep"
	"testing"
)

func TestViaCepIntegration(t *testing.T) {
	t.Run("SUCESS: Get a address by cep", func(t *testing.T) {
		zipcode := cep.CEP("01001-000")
		integration := viacep.NewIntegration(zipcode)

		expectedZipcode := zipcode
		expectedAddress := "Praça da Sé"

		address, err := integration.GetAddress()
		if err != nil {
			t.Fatal("Errors was not expected!", err.Error())
		}

		if address.Logradouro != expectedAddress {
			t.Errorf("Was expected '%s', but got '%s'", expectedAddress, address.Logradouro)
		}

		if address.CEP != expectedZipcode {
			t.Errorf("Was expected '%s', but got '%s'", expectedZipcode, address.CEP)
		}
	})

	t.Run("ERROR: Invalid zipcode cause a bad request", func(t *testing.T) {
		zipcode := cep.CEP("01001")
		integration := viacep.NewIntegration(zipcode)

		_, err := integration.GetAddress()
		if err == nil {
			t.Fatal("Errors was expected!")
		}
	})
}
