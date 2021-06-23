package gopspa

type Emprestimo struct {
	Id    int   `json:"id"`
	Aluno Aluno `json:"aluno"`
	Livro Livro `json:"livro"`
}
