[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

# Go ViaCep
##### This lib was developed for get an adress data by zipcode using [ViaCep](https://viacep.com.br/) API

## Import and Usage
```
go get github.com/WalterPaes/go-viacep
```

### Create a New Integration
```go
// Parse zipcode to CEP type
zipcode := cep.CEP("01001-000")

// Create a new integration
integration := viacep.NewIntegration(zipcode)
```

### Get Address
```go
integration.GetAddress()
```

## :rocket: Technologies

This project was developed with the following technologies:

-  [Go](https://golang.org/)
-  [GoLand](https://www.jetbrains.com/go/?gclid=EAIaIQobChMI5-ug_OvG6gIVBgiRCh0GGARZEAAYASAAEgKOSPD_BwE)
-  [ViaCep](https://viacep.com.br/)

Made by [Walter Junior](https://www.linkedin.com/in/walter-paes/)
