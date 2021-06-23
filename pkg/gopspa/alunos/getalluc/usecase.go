package getalluc

import "gopspa/pkg/gopspa"

type usecase struct {
	Repo AlunosRepo
}

func New(repo AlunosRepo) *usecase {
	return &usecase{
		Repo: repo,
	}
}

func (uc *usecase) Run() (alunos []gopspa.Aluno, err error) {
	alunos, err = uc.Repo.FindAll()

	if err != nil {
		return
	}

	return
}
