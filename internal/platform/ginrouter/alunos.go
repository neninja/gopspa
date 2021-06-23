package ginrouter

import (
	"gopspa/internal/platform/inmemo"
	"gopspa/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Alunos(r *gin.RouterGroup) {
	r.GET("/alunos", gin.WrapF(getallalunos))
	r.POST("/alunos", gin.WrapF(createaluno))
}

// @description Lista de todos Alunos
// @tags alunos
// @produce json
// @success 200 {array} web.GetAllAlunosOutput
// @failure 500 {object} web.RequestError{} "Generic error"
// @router	/api/alunos [get]
func getallalunos(w http.ResponseWriter, req *http.Request) {
	repo := inmemo.AlunosRepo{}
	web.GetAllAlunos(repo)(w, req)
}

// @description Cria um aluno
// @tags alunos
// @accept json
// @produce json
// @param input body web.CreateAlunoInput true "Aluno"
// @success 200 {object} gopspa.Aluno
// @failure 500 {object} web.RequestError{} "Generic error"
// @router	/api/aluno [post]
func createaluno(w http.ResponseWriter, req *http.Request) {
	r := inmemo.AlunosRepo{}
	web.CreateAluno(r)(w, req)
}
