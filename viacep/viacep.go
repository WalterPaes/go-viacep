package viacep

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

const (
	cepNotFoundError = "CEP '%s' was not found"
	invalidCepError = "'%s' is an invalid CEP"
)

type Address struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	Unidade     string `json:"unidade,omitempty"`
	IBGE        string `json:"ibge,omitempty"`
	Gia         string `json:"gia,omitempty"`
}

type ViaCep struct {
	cep string
}

func NewViaCEP(cep string) (*ViaCep, error) {
	if err := validate(cep); err != nil {
		log.Printf("[VIACEP] VALIDATION ERROR: %s \n", err.Error())
		return nil, err
	}

	return &ViaCep{
		cep: cep,
	}, nil
}

func (vc ViaCep) GetAddress() (*Address, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%v/json/", vc.cep)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("[VIACEP] REQUEST ERROR: %s \n", err.Error())
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		notFoundError := fmt.Errorf(cepNotFoundError, vc.cep)
		log.Printf("[VIACEP] NOT FOUND ERROR: " + notFoundError.Error())
		return nil, notFoundError
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[VIACEP] BODY PARSE ERROR: %s \n", err.Error())
		return nil, err
	}

	var address *Address
	err = json.Unmarshal(body, &address)
	if err != nil {
		log.Printf("[VIACEP] JSON PARSE ERROR: %s \n", err.Error())
		return nil, err
	}

	return address, err
}

func validate(cep string) error {
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		return err
	}

	cep = reg.ReplaceAllString(cep, "")
	if len(cep) < 7 || len(cep) > 8 {
		return fmt.Errorf(invalidCepError, cep)
	}
	return nil
}