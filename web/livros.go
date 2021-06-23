package web

import (
	"encoding/json"
	"gopspa/pkg/gopspa"
	"gopspa/pkg/gopspa/livros/createuc"
	"gopspa/pkg/gopspa/livros/getalluc"
	"net/http"
)

type CreateLivroInput struct {
	Nome string `json:"nome"`
	ISBN string `json:"isbn"`
}

func CreateLivro(r createuc.LivrosRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		uc := createuc.New(r)

		var input CreateLivroInput
		err := json.NewDecoder(req.Body).Decode(&input)
		if err != nil {
			JSONError(w, 500, err.Error())
			return
		}

		livro, err := uc.Run(input.Nome, input.ISBN)

		if err != nil {
			JSONError(w, 500, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(livro)

		if err != nil {
			JSONError(w, 500, err.Error())
			return
		}
	}
}

type GetAllLivroOutput struct {
	Livros []gopspa.Livro `json:"livros"`
}

func GetAllLivros(repo getalluc.LivrosRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		uc := getalluc.New(repo)
		livros, err := uc.Run()

		if err != nil {
			JSONError(w, 500, err.Error())
			return
		}

		output := GetAllLivroOutput{livros}

		// no frontend é preferivel voltar um array vazio
		if livros == nil {
			output.Livros = make([]gopspa.Livro, 0)
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(output)

		if err != nil {
			JSONError(w, 500, err.Error())
			return
		}
	}
}
