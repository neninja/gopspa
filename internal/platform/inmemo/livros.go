package inmemo

import (
	"errors"
	"gopspa/pkg/gopspa"
)

type LivrosRepo struct{}

var lper []gopspa.Livro

func (r LivrosRepo) FindAll() ([]gopspa.Livro, error) {
	return lper, nil
}

func (r LivrosRepo) Find(id int) (l gopspa.Livro, err error) {
	for _, li := range lper {
		if id == li.Id {
			l = li
			return
		}
	}
	err = errors.New("Livro não encontrado")
	return
}

func (r LivrosRepo) Save(l gopspa.Livro) (gopspa.Livro, error) {
	id := len(lper)
	l.Id = id
	lper = append(lper, l)
	return l, nil
}
