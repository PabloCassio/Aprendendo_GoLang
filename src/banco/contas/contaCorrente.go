package contas

import "banco/clientes"

type ContaCorrente struct { //estrutura
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo = c.saldo - valorDoSaque
		return "Saque efetuado com sucesso"
	}
	return "saldo insuficiente"
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) { // * recebe endereço passado por &
	if valorDoDeposito > 0 {
		c.saldo = c.saldo + valorDoDeposito
		return "Depósito efetuado com sucesso. Seu saldo agora é ", c.saldo

	}
	return "Valor do depósito incorreto. Seu saldo atual é", c.saldo

}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia <= c.saldo || valorTransferencia > 0 {
		c.saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	}
	return false

}

func (c *ContaCorrente) GetSaldo() float64 {
	return c.saldo
}
