package createuc

import "gopspa/pkg/gopspa"

type usecase struct {
	Repo LivrosRepo
}

func New(r LivrosRepo) *usecase {
	return &usecase{
		Repo: r,
	}
}

func (uc *usecase) Run(nome string, isbn string) (livro gopspa.Livro, err error) {
	livro.Nome = nome
	livro.ISBN = isbn

	livro, err = uc.Repo.Save(livro)
	if err != nil {
		return
	}

	return
}
