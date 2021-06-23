package createuc

import "gopspa/pkg/gopspa"

type LivrosRepo interface {
	Find(id int) (gopspa.Livro, error)
}

type AlunosRepo interface {
	Find(id int) (gopspa.Aluno, error)
}

type EmprestimosRepo interface {
	Save(emprestimo gopspa.Emprestimo) (gopspa.Emprestimo, error)
}
