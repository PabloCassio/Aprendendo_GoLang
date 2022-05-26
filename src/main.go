package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta VerificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type VerificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoPablo := contas.ContaCorrente{}
	contaDoPablo.Depositar(500)

	contaDoMario := contas.ContaPoupanca{}
	contaDoMario.Depositar(200)

	PagarBoleto(&contaDoMario, 300)
	PagarBoleto(&contaDoPablo, 300)
	fmt.Println(contaDoMario.GetSaldo())
	fmt.Println(contaDoPablo.GetSaldo())
}
