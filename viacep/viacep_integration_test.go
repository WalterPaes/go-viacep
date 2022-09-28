package viacep

import (
	"fmt"
	"testing"
)

func TestViaCepIntegration(t *testing.T) {
	t.Run("SUCESS: Get a address by cep", func(t *testing.T) {
		cep := "01001-000"
		viaCep, err := NewViaCEP(cep)
		if err != nil {
			t.Fatal("Errors was not expected!", err.Error())
		}
		
		address, err := viaCep.GetAddress()
		if err != nil {
			t.Fatal("Errors was not expected!", err.Error())
		}

		if address.CEP != cep {
			t.Errorf("Was expected '%s', but got '%s'", cep, address.CEP)
		}

		expectedAddress := Address{
			Logradouro: "Praça da Sé",
			Complemento: "lado ímpar",
			Bairro: 	"Sé",
			Localidade: "São Paulo",
			UF: 	"SP",
		}

		if address.Logradouro != expectedAddress.Logradouro {
			t.Errorf("Was expected '%s', but got '%s'", expectedAddress.Logradouro, address.Logradouro)
		}

		if address.Complemento != expectedAddress.Complemento {
			t.Errorf("Was expected '%s', but got '%s'", expectedAddress.Complemento, address.Complemento)
		}

		if address.Bairro != expectedAddress.Bairro {
			t.Errorf("Was expected '%s', but got '%s'", expectedAddress.Bairro, address.Bairro)
		}

		if address.Localidade != expectedAddress.Localidade {
			t.Errorf("Was expected '%s', but got '%s'", expectedAddress.Localidade, address.Localidade)
		}

		if address.UF != expectedAddress.UF {
			t.Errorf("Was expected '%s', but got '%s'", expectedAddress.UF, address.UF)
		}
	})

	t.Run("ERROR: Not found zipcode", func(t *testing.T) {
		cep := "1234567"
		viaCep, err := NewViaCEP(cep)
		if err != nil {
			t.Fatal("Errors was not expected!", err.Error())
		}

		_, err = viaCep.GetAddress()
		notFoundError := fmt.Errorf(cepNotFoundError, cep)
		if err.Error() != notFoundError.Error() {
			t.Errorf("Was expected '%s', but got '%s'", notFoundError.Error(), err.Error())
		}
	})

	t.Run("ERROR: Invalid zipcode", func(t *testing.T) {
		cep := "12345"
		_, err := NewViaCEP(cep)
		
		invalidError := fmt.Errorf(invalidCepError, cep)
		if err.Error() != invalidError.Error() {
			t.Errorf("Was expected '%s', but got '%s'", invalidError.Error(), err.Error())
		}
	})
}
