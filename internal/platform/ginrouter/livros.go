package ginrouter

import (
	"gopspa/internal/platform/inmemo"
	"gopspa/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Livros(r *gin.RouterGroup) {
	r.GET("/livros", gin.WrapF(getalllivros))
	r.POST("/livros", gin.WrapF(createlivro))
}

// @description Lista de todos livros
// @tags livros
// @produce json
// @success 200 {array} gopspa.Livro
// @failure 500 {object} web.RequestError{} "Generic error"
// @router	/api/livros [get]
func getalllivros(w http.ResponseWriter, req *http.Request) {
	repo := inmemo.LivrosRepo{}
	web.GetAllLivros(repo)(w, req)
}

// @description Cria um livro
// @tags livros
// @accept json
// @produce json
// @param input body web.CreateLivroInput true "Livro"
// @success 200 {object} web.GetAllLivroOutput
// @failure 500 {object} web.RequestError{} "Generic error"
// @router	/api/livros [post]
func createlivro(w http.ResponseWriter, req *http.Request) {
	r := inmemo.LivrosRepo{}
	web.CreateLivro(r)(w, req)
}
