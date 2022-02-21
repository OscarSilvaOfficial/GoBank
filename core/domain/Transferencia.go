package domain

import (
	"fmt"

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

func (transferencia *Transferencia) Transferir(valor float64) {
	fmt.Println("\n>>>>>>> TRANSFERÊNCIA")
	fmt.Printf("Saldo da conta que fará a transferência R$%.2f\n", transferencia.contaOrigem.Saldo())
	fmt.Printf("Saldo da conta que receberá a transferência R$%.2f\n", transferencia.contaDestino.Saldo())
	transferencia.contaOrigem.Sacar(valor)
	transferencia.contaDestino.Depositar(valor)
	fmt.Printf("Saldo da conta que fez a transferência R$%.2f\n", transferencia.contaOrigem.Saldo())
	fmt.Printf("Saldo da conta que recebeu a transferência R$%.2f\n", transferencia.contaDestino.Saldo())
}
