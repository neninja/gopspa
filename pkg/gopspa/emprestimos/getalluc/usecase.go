package getalluc

import (
	"gopspa/pkg/gopspa"
)

type usecase struct {
	Repo EmprestimosRepo
}

func New(r EmprestimosRepo) *usecase {
	return &usecase{
		Repo: r,
	}
}

func (uc *usecase) Run() (emprestimos []gopspa.Emprestimo, err error) {
	emprestimos, err = uc.Repo.FindAll()
	if err != nil {
		return
	}

	return
}
