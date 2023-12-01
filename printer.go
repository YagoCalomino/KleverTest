package main

import (
	"encoding/json"
	"fmt"
)

func PrintAddressDetails(response []byte) {
	var data map[string]interface{}
	err := json.Unmarshal(response, &data)
	if err != nil {
		fmt.Println("Erro ao fazer o parse da resposta:", err)
		return
	}

	fmt.Println("Address:", data["address"])
	fmt.Println("Balance:", data["balance"])
	fmt.Println("Total Transactions:", data["totalTx"])
}

func PrintTransactionDetails(response []byte) {
	var data map[string]interface{}
	err := json.Unmarshal(response, &data)
	if err != nil {
		fmt.Println("Erro ao fazer o parse da resposta:", err)
		return
	}

	fmt.Println("Transaction ID:", data["txID"])
	fmt.Println("Block:", data["block"])
	// Exibir os endereços e valores envolvidos na transação
	fmt.Println("Addresses and Values:")
	addresses := data["addresses"].([]interface{})
	for _, addr := range addresses {
		addressData := addr.(map[string]interface{})
		fmt.Printf("Address: %s, Value: %s\n", addressData["address"], addressData["value"])
	}
}
