package port

type ContaContrato interface {
	Sacar(valor float64) float64
	Depositar(valor float64) float64
	Saldo() float64
}

