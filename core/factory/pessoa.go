package factory

import "github.com/OscarSilvaOfficial/GoBank/core/domain"

func FactoryPessoa(nome string, documento string) domain.Pessoa {
	return domain.Pessoa{Nome: nome, Documento: documento}
}
