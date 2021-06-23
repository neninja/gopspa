package ginrouter

import (
	"gopspa/internal/platform/inmemo"
	"gopspa/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Emprestimos(r *gin.RouterGroup) {
	r.GET("/emprestimos", gin.WrapF(getallemprestimos))
	r.POST("/emprestimos", gin.WrapF(createemprestimo))
}

// @description Lista de todos emprestimos
// @tags emprestimos
// @produce json
// @success 200 {array} web.GetAllEmprestimos
// @failure 500 {object} web.RequestError{} "Generic error"
// @router /api/enprestimos [get]
func getallemprestimos(w http.ResponseWriter, req *http.Request) {
	repo := inmemo.EmprestimosRepo{}
	web.GetAllEmprestimos(repo)(w, req)
}

// @description Lista de todos emprestimos
// @tags emprestimos
// @accept json
// @produce json
// @param input body web.CreateEmprestimoInput true "Emprestimo"
// @success 200 {object} gopspa.Emprestimo
// @failure 500 {object} web.RequestError{} "Generic error"
// @router /api/emprestimos [post]
func createemprestimo(w http.ResponseWriter, req *http.Request) {
	r1 := inmemo.AlunosRepo{}
	r2 := inmemo.LivrosRepo{}
	r3 := inmemo.EmprestimosRepo{}
	web.CreateEmprestimo(r1, r2, r3)(w, req)
}
