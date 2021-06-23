package createuc

import "gopspa/pkg/gopspa"

type LivrosRepo interface {
	Save(emprestimo gopspa.Livro) (gopspa.Livro, error)
}
