package contas

import "banco/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo = c.saldo - valorDoSaque
		return "Saque efetuado com sucesso"
	}
	return "saldo insuficiente"
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo = c.saldo + valorDoDeposito
		return "Depósito efetuado com sucesso. Seu saldo agora é ", c.saldo

	}
	return "Valor do depósito incorreto. Seu saldo atual é", c.saldo

}

func (c *ContaPoupanca) GetSaldo() float64 {
	return c.saldo
}
