package web

import (
	"encoding/json"
	"gopspa/pkg/gopspa"
	"gopspa/pkg/gopspa/alunos/createuc"
	"gopspa/pkg/gopspa/alunos/getalluc"
	"net/http"
)

type CreateAlunoInput struct {
	Nome string `json:"nome"`
}

func CreateAluno(r createuc.AlunosRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		uc := createuc.New(r)

		var input CreateAlunoInput
		err := json.NewDecoder(req.Body).Decode(&input)
		if err != nil {
			JSONError(w, 500, err.Error())
			return
		}

		aluno, err := uc.Run(input.Nome)

		if err != nil {
			JSONError(w, 500, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(aluno)

		if err != nil {
			JSONError(w, 500, err.Error())
			return
		}
	}
}

type GetAllAlunosOutput struct {
	Alunos []gopspa.Aluno `json:"alunos"`
}

func GetAllAlunos(repo getalluc.AlunosRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		uc := getalluc.New(repo)
		alunos, err := uc.Run()

		if err != nil {
			JSONError(w, 500, err.Error())
			return
		}

		output := GetAllAlunosOutput{alunos}

		// no frontend é preferivel voltar um array vazio
		if alunos == nil {
			output.Alunos = make([]gopspa.Aluno, 0)
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(output)

		if err != nil {
			JSONError(w, 500, err.Error())
			return
		}
	}
}
