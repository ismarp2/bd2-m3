package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("livro", repository_livro.AdicionarLivro)
	router.GET("livro", repository_livro.ObterLivros)
	router.PUT("livro", repository_livro.AtualizarLivro)
	router.DELETE("livro", repository_livro.RemoverLivro)

	router.POST("usuario", repository_usuario.AdicionarUsuario)
	router.GET("usuario", repository_usuario.ObterUsuarios)
	router.PUT("usuario", repository_usuario.AtualizarUsuario)
	router.DELETE("usuario", repository_usuario.RemoverUsuario)

	router.POST("locacao", repository_locacao.AdicionarUsuario)
	router.GET("locacao", repository_locacao.ObterUsuarios)
	router.PUT("locacao", repository_locacao.AtualizarUsuario)
	router.DELETE("locacao", repository_locacao.RemoverUsuario)

	router.Run("localhost:8080")
}
