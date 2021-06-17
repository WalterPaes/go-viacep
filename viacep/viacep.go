package viacep

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/WalterPaes/go-viacep/cep"
	"io/ioutil"
	"log"
	"net/http"
)

// Integration represents viacep integration
type Integration struct {
	cep cep.CEP
}

func NewIntegration(cep cep.CEP) *Integration {
	return &Integration{cep: cep}
}

func (vcs Integration) GetAddress() (*cep.Address, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%v/json/", vcs.cep)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("[VIACEP] ERROR: %s \n", err.Error())
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("[VIACEP] ERROR: " + resp.Status)
		return nil, errors.New(resp.Status)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[VIACEP] ERROR: %s \n", err.Error())
		return nil, err
	}

	var address *cep.Address
	err = json.Unmarshal(body, &address)
	if err != nil {
		log.Printf("[VIACEP] ERROR: %s \n", err.Error())
		return nil, err
	}

	return address, err
}
