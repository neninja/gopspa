package getalluc

import "gopspa/pkg/gopspa"

type AlunosRepo interface {
	FindAll() ([]gopspa.Aluno, error)
}
