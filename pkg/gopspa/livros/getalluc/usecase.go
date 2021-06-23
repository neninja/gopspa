package getalluc

import "gopspa/pkg/gopspa"

type usecase struct {
	Repo LivrosRepo
}

func New(repo LivrosRepo) *usecase {
	return &usecase{
		Repo: repo,
	}
}

func (uc *usecase) Run() (livros []gopspa.Livro, err error) {
	livros, err = uc.Repo.FindAll()

	if err != nil {
		return
	}

	return
}
