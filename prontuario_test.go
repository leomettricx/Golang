/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const (
	NOME       = "Nome"
	ENDERECO   = "Endereco"
	CEP        = "Cep"
	CIDADE     = "Cidade"
	ESTADO     = "Estado"
	TELEFONE1  = "Telefone1"
	TELEFONE2  = "Telefone2"
	EMAIL      = "Email"
	NASCIMENTO = "Nascimento"
	CPF        = "CPF"
)

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func getInitArguments() [][]byte {
	return [][]byte{[]byte("init"),
		[]byte("Leonardo Moreira de Oliveira"),
		[]byte("Rua Sto Amaro, 200/230. Gloria. RJ"),
		[]byte("22211-230"),
		[]byte("Rio de janeiro"),
		[]byte("RJ"),
		[]byte("5521994073374"),
		[]byte("5521994073374"),
		[]byte("leo@mettricx.com"),
		[]byte("06/12/1961"),
		[]byte("765.476.365.555-72"),
		[]byte("A+")}
}

func TestProntuario_Init(t *testing.T) {
	scc := new(ProntuarioChaincode)
	scc.testMode = true
	stub := shim.NewMockStub("Prontuario", scc)

	// Init
	checkInit(t, stub, getInitArguments())

	checkState(t, stub, "Nome", NOME)
	checkState(t, stub, "Endereco", ENDERECO)
	checkState(t, stub, "Cep", CEP)
	checkState(t, stub, "Cidade", CIDADE)
	checkState(t, stub, "Estado", ESTADO)
	checkState(t, stub, "Telefone1", TELEFONE1)
	checkState(t, stub, "Telefone2", TELEFONE2)
	checkState(t, stub, "Email", EMAIL)
	checkState(t, stub, "Nascimento", NASCIMENTO)
	checkState(t, stub, "Cpf", CPF)
}

func checkState(t *testing.T, stub *shim.MockStub, name string, value string) {
	bytes := stub.State[name]
	if bytes == nil {
		fmt.Println("State", name, "failed to get value")
		t.FailNow()
	}
	if string(bytes) != value {
		fmt.Println("State value", name, "was", string(bytes), "and not", value, "as expected")
		t.FailNow()
	}
}
