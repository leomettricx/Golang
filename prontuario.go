/*
Copyright Mettricx. 2019 All Rights Reserved.

*/

package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// ProntuarioChaincode implementation
type ProntuarioChaincode struct {
	testMode bool
}

func (t *ProntuarioChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Criando Prontuario")
	// args é vetor que recebe todos os parâmetros de entrada
	_, args := stub.GetFunctionAndParameters()
	var err error

	// Upgrade Mode 1: leave ledger state as it was
	if len(args) == 0 {
		return shim.Success(nil)
	}

	fmt.Printf("Nome      : %s\n", args[0])
	fmt.Printf("Endereco  : %s\n", args[1])
	fmt.Printf("Cep       : %s\n", args[2])
	fmt.Printf("Cidade    : %s\n", args[3])
	fmt.Printf("Estado    : %s\n", args[4])
	fmt.Printf("Telefone1 : %s\n", args[5])
	fmt.Printf("Telefone2 : %s\n", args[6])
	fmt.Printf("Nascimento: %s\n", args[7])
	fmt.Printf("Email     : %s\n", args[8])
	fmt.Printf("CPF	      : %s\n", args[9])
	fmt.Printf("TipoSanguineo     : %s\n", args[10])

	// Map participant identities to their roles on the ledger
	roleKeys := []string{nomKey, endKey, cepKey, cidKey, estKey, te1Key, te2Key, nasKey, emaKey, cpfKey, tsgKey}
	for i, roleKey := range roleKeys {
		err = stub.PutState(roleKey, []byte(args[i]))
		if err != nil {
			fmt.Errorf("Error recording key %s: %s\n", roleKey, err.Error())
			return shim.Error(err.Error())
		}
	}

	return shim.Success(nil)
}

func (t *ProntuarioChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("TradeWorkflow Invoke")

	return shim.Error("Invalid invoke function name")
}
