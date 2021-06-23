package createuc

import "gopspa/pkg/gopspa"

type AlunosRepo interface {
	Save(emprestimo gopspa.Aluno) (gopspa.Aluno, error)
}
