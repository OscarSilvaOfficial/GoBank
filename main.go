package main

import (
	"fmt"

	"github.com/OscarSilvaOfficial/GoBank/core/domain"
	"github.com/OscarSilvaOfficial/GoBank/core/factory"
)

func main() {
	contaCorrenteOscar := factory.FactoryContaCorrente("Oscar", "07244444444", 000, 1111103, 1000.0)
	contaCorrenteDondi := factory.FactoryContaCorrente("Dondi", "07244444441", 000, 12321309, 1000.0)
	contaCorrenteBla := factory.FactoryContaCorrente("Bla", "07244444443", 001, 03210312, 1000.0)

	todasAsContas := []domain.ContaCorrente{
		contaCorrenteOscar,
		contaCorrenteDondi,
		contaCorrenteBla,
	}

	MostrarContas(todasAsContas)

	fmt.Println("\n>>>>>>> SAC")
	fmt.Printf("Conta Oscar saldo atual R$%.2f\n", contaCorrenteOscar.Saldo())
	contaCorrenteOscar.Sacar(1000.1) // Saldo insuficiente
	contaCorrenteOscar.Sacar(-1)     // Não é possivel sacar valores negativos
	contaCorrenteOscar.Sacar(2)      // Saque efetuado
	fmt.Printf("Conta Oscar saldo atual R$%.2f\n", contaCorrenteOscar.Saldo())

	fmt.Println("\n>>>>>>> DEPOSITO")
	fmt.Printf("Conta Oscar saldo atual R$%.2f\n", contaCorrenteOscar.Saldo())
	contaCorrenteOscar.Depositar(-1) // Não é possivel depositar valores negativos
	contaCorrenteOscar.Depositar(2)  // Deposito realizado
	fmt.Printf("Conta Oscar saldo atual R$%.2f\n", contaCorrenteOscar.Saldo())

	fmt.Println("\n>>>>>>> TRANSFERÊNCIA")
	fmt.Printf("Saldo da conta que fará a transferência R$%.2f\n", contaCorrenteOscar.Saldo())
	fmt.Printf("Saldo da conta que receberá a transferência R$%.2f\n", contaCorrenteDondi.Saldo())
	transferencia := factory.FactoryTransferencia(&contaCorrenteOscar, &contaCorrenteDondi, 100.0)
	fmt.Printf("Saldo da conta que fez a transferência R$%.2f\n", transferencia.ContaOrigem().Saldo())
	fmt.Printf("Saldo da conta que recebeu a transferência R$%.2f\n", transferencia.ContaDestino().Saldo())

}

func MostrarContas(contas []domain.ContaCorrente) {
	fmt.Println("\n>>>>>>> CONTAS ABERTAS")
	for _, conta := range contas {
		fmt.Printf("Conta %s: ", conta.Titular().Nome)
		fmt.Println(conta)
	}
}
