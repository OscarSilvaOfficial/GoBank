package domain

import (
	"github.com/OscarSilvaOfficial/GoBank/port"
)

type Transferencia struct {
	contaOrigem  port.ContaContrato
	contaDestino port.ContaContrato
}

func (t *Transferencia) ContaOrigem() port.ContaContrato {
	return t.contaOrigem
}

func (t *Transferencia) ContaDestino() port.ContaContrato {
	return t.contaDestino
}

func (trenferencia *Transferencia) SetContaOrigem(contaOrigem port.ContaContrato) *Transferencia {
	trenferencia.contaOrigem = contaOrigem
	return trenferencia
}

func (trenferencia *Transferencia) SetContaDestino(contaDestino port.ContaContrato) *Transferencia {
	trenferencia.contaDestino = contaDestino
	return trenferencia
}

func (transferencia *Transferencia) Transferir(valor float64) Transferencia {
	transferencia.contaOrigem.Sacar(valor)
	transferencia.contaDestino.Depositar(valor)
	return *transferencia
}
