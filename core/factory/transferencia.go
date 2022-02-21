package factory

import (
	"github.com/OscarSilvaOfficial/GoBank/core/domain"
	"github.com/OscarSilvaOfficial/GoBank/port"
)

func FactoryTransferencia(contaOrigem, contaDestino port.ContaContrato, valor float64) {
	novaTransferencia := domain.Transferencia{}
	novaTransferencia.SetContaOrigem(contaOrigem).SetContaDestino(contaDestino).Transferir(valor)
}
