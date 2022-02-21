package factory

import (
	"github.com/OscarSilvaOfficial/GoBank/core/domain"
)

func FactoryContaCorrente(titular string, documento string, agencia int, conta int, saldo float64) domain.ContaCorrente {
	novaPessoa := FactoryPessoa(titular, documento)
	novaConta := domain.ContaCorrente{}
	novaConta.SetTitular(novaPessoa).SetAgencia(agencia).SetConta(conta).SetSaldo(saldo)
	return novaConta
}
