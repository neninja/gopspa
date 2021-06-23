package web

import (
	"encoding/json"
	"gopspa/pkg/gopspa"
	"gopspa/pkg/gopspa/emprestimos/createuc"
	"gopspa/pkg/gopspa/emprestimos/getalluc"
	"net/http"
)

type CreateEmprestimoInput struct {
	IdAluno int `json:"idAluno"`
	IdLivro int `json:"idLivro"`
}

func CreateEmprestimo(r1 createuc.AlunosRepo, r2 createuc.LivrosRepo, r3 createuc.EmprestimosRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		uc := createuc.New(r3, r1, r2)

		var input CreateEmprestimoInput
		err := json.NewDecoder(req.Body).Decode(&input)
		if err != nil {
			JSONError(w, 500, err.Error())
			return
		}

		emprestimo, err := uc.Run(input.IdAluno, input.IdLivro)

		if err != nil {
			JSONError(w, 500, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(emprestimo)

		if err != nil {
			JSONError(w, 500, err.Error())
			return
		}
	}
}

type GetAllEmprestimosOutput struct {
	Emprestimos []gopspa.Emprestimo `json:"emprestimos"`
}

func GetAllEmprestimos(repo getalluc.EmprestimosRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		uc := getalluc.New(repo)
		emprestimos, err := uc.Run()

		if err != nil {
			JSONError(w, 500, err.Error())
			return
		}

		output := GetAllEmprestimosOutput{emprestimos}

		// no frontend é preferivel voltar um array vazio
		if emprestimos == nil {
			output.Emprestimos = make([]gopspa.Emprestimo, 0)
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(output)

		if err != nil {
			JSONError(w, 500, err.Error())
			return
		}
	}
}
