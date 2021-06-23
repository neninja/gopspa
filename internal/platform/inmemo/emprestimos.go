package inmemo

import (
	"gopspa/pkg/gopspa"
)

type EmprestimosRepo struct{}

var eper []gopspa.Emprestimo

func (r EmprestimosRepo) Save(e gopspa.Emprestimo) (gopspa.Emprestimo, error) {
	id := len(eper)
	e.Id = id
	eper = append(eper, e)
	return e, nil
}

func (r EmprestimosRepo) FindAll() ([]gopspa.Emprestimo, error) {
	return eper, nil
}
