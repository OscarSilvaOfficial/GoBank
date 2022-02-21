package domain

import "fmt"

type ContaCorrente struct {
	titular        Pessoa
	agencia, conta int
	saldo          float64
}

func (conta *ContaCorrente) ValidaSac(valor float64) bool {
	if conta.saldo < valor {
		fmt.Println("Saldo insuficiente")
		return false
	}

	if valor < 0 {
		fmt.Println("Não e possivel sacar valores negatigos")
		return false
	}

	return true
}

func (conta *ContaCorrente) ValidaDeposito(valor float64) bool {
	if valor < 0 {
		fmt.Println("Não e possivel depositar valores negativos")
		return false
	}
	return true
}

func (conta *ContaCorrente) Sacar(valor float64) float64 {
	if conta.ValidaSac(valor) {
		conta.saldo -= valor
	}
	return conta.saldo
}

func (conta *ContaCorrente) Depositar(valor float64) float64 {
	if conta.ValidaDeposito(valor) {
		conta.saldo += valor
	}
	return conta.saldo
}

func (conta *ContaCorrente) Saldo() float64 {
	return conta.saldo
}

func (conta *ContaCorrente) SetTitular(pessoa Pessoa) *ContaCorrente {
	conta.titular = pessoa
	return conta
}

func (conta *ContaCorrente) SetAgencia(agencia int) *ContaCorrente {
	conta.agencia = agencia
	return conta
}

func (conta *ContaCorrente) SetConta(Nconta int) *ContaCorrente {
	conta.conta = Nconta
	return conta
}

func (conta *ContaCorrente) SetSaldo(saldo float64) *ContaCorrente {
	conta.saldo = saldo
	return conta
}
