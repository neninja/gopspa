package createuc

import "gopspa/pkg/gopspa"

type usecase struct {
	Repo AlunosRepo
}

func New(r AlunosRepo) *usecase {
	return &usecase{
		Repo: r,
	}
}

func (uc *usecase) Run(nome string) (aluno gopspa.Aluno, err error) {
	aluno.Nome = nome

	aluno, err = uc.Repo.Save(aluno)
	if err != nil {
		return
	}

	return
}
