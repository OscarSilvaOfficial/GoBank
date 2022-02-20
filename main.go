package main

import (
	"fmt"

	"github.com/OscarSilvaOfficial/GoBank/factory"
)

func main() {
	contaCorrenteOscar := factory.FactoryContaCorrente("Oscar", "07244444444", 000, 1111103, 1000.0)
	contaCorrenteDondi := factory.FactoryContaCorrente("Dondi", "07244444441", 000, 12321309, 1)
	contaCorrentBla := factory.FactoryContaCorrente("Bla", "07244444443", 001, 03210312, 0)

	fmt.Println("\n>>>>>>> CONTAS ABERTAS")
	fmt.Println("Conta Oscar", contaCorrenteOscar)
	fmt.Println("Conta Dondi", contaCorrenteDondi)
	fmt.Println("Conta Bla", contaCorrentBla)

	fmt.Println("\n>>>>>>> SAC")
	contaCorrenteOscar.Sacar(1000.1) // Saldo insuficiente
	contaCorrenteOscar.Sacar(-1)     // Não é possivel sacar valores negativos
	contaCorrenteOscar.Sacar(2)      // Saque efetuado
	fmt.Println("Conta Oscar", contaCorrenteOscar)

	fmt.Println("\n>>>>>>> DEPOSITO")
	contaCorrenteOscar.Depositar(-1) // Não é possivel depositar valores negativos
	contaCorrenteOscar.Depositar(2)  // Deposito realizado
	fmt.Println("Conta Oscar", contaCorrenteOscar)

}
