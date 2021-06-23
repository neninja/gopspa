package createuc

import (
	"errors"
	"gopspa/pkg/gopspa"
)

type usecase struct {
	EmpRepo    EmprestimosRepo
	AlunosRepo AlunosRepo
	LivrosRepo LivrosRepo
}

func New(erepo EmprestimosRepo, arepo AlunosRepo, lrepo LivrosRepo) *usecase {
	return &usecase{
		EmpRepo:    erepo,
		AlunosRepo: arepo,
		LivrosRepo: lrepo,
	}
}

func (uc *usecase) Run(idAluno int, idLivro int) (emprestimo gopspa.Emprestimo, err error) {
	aluno, err := uc.AlunosRepo.Find(idAluno)
	if err != nil {
		return
	}
	if aluno == (gopspa.Aluno{}) {
		err = errors.New("Usuário não localizado")
	}

	livro, err := uc.LivrosRepo.Find(idAluno)
	if err != nil {
		return
	}
	if livro == (gopspa.Livro{}) {
		err = errors.New("Livro não localizado")
	}

	emprestimo = gopspa.Emprestimo{Aluno: aluno, Livro: livro}
	emprestimo, err = uc.EmpRepo.Save(emprestimo)
	if err != nil {
		return
	}

	return
}
