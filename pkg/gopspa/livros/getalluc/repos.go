package getalluc

import "gopspa/pkg/gopspa"

type LivrosRepo interface {
	FindAll() ([]gopspa.Livro, error)
}
