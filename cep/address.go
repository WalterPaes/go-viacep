package cep

type CEP string

type Address struct {
	CEP         `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	Unidade     string `json:"unidade,omitempty"`
	IBGE        string `json:"ibge,omitempty"`
	Gia         string `json:"gia,omitempty"`
}
