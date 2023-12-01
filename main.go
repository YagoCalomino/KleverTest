package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	address := "bc1qyzxdu4px4jy8gwhcj82zpv7qzhvc0fvumgnh0r"
	txID := "3654d26660dcc05d4cfb25a1641a1e61f06dfeb38ee2279bdb049d018f1830ab"

	addressDetailsEndpoint := fmt.Sprintf("https://bitcoin.explorer.klever.io/api/v2/address/%s", address)
	addressResponse := sendRequest(addressDetailsEndpoint)
	printResponse(addressResponse)

	txDetailsEndpoint := fmt.Sprintf("https://bitcoin.explorer.klever.io/api/v2/tx/%s", txID)
	txResponse := sendRequest(txDetailsEndpoint)
	printResponse(txResponse)

	PrintAddressDetails(addressResponse)
	PrintTransactionDetails(txResponse)
}

func sendRequest(endpoint string) []byte {
	username := "support"
	password := "Fg+GJKDACKIEOD3XVps="

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Println("Erro ao criar a solicitação:", err)
		return nil
	}

	req.SetBasicAuth(username, password)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao enviar a solicitação:", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler a resposta:", err)
		return nil
	}

	return body
}

func printResponse(response []byte) {
	var data interface{}
	err := json.Unmarshal(response, &data)
	if err != nil {
		fmt.Println("Erro ao fazer o parse da resposta:", err)
		return
	}

	prettyJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Erro ao formatar a resposta:", err)
		return
	}

	fmt.Println(string(prettyJSON))
}
