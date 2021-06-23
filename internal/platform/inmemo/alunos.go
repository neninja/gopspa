package inmemo

import (
	"errors"
	"gopspa/pkg/gopspa"
)

type AlunosRepo struct{}

var aper []gopspa.Aluno

func (r AlunosRepo) FindAll() ([]gopspa.Aluno, error) {
	return aper, nil
}

func (r AlunosRepo) Find(id int) (a gopspa.Aluno, err error) {
	for _, ai := range aper {
		if id == ai.Id {
			a = ai
			return
		}
	}
	err = errors.New("Aluno não encontrado")
	return
}

func (r AlunosRepo) Save(a gopspa.Aluno) (gopspa.Aluno, error) {
	id := len(aper)
	a.Id = id
	aper = append(aper, a)
	return a, nil
}
