package repository_livro

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func GetAlbums(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, albums)
// }

func ObterLivros(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func AdicionarLivro(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func AtualizarLivro(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func RemoverLivro(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
