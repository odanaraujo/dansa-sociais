package rotas

import (
	"net/http"

	"github.com/dansa-sociais/api/rsc/controllers"
)

var rotasUsuarios = []Rota{
	{
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},

	{
		URI:                "/usuarios/",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarTodosUsuarios,
		RequerAutenticacao: false,
	},

	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.AtualizarUsuario,
		RequerAutenticacao: false,
	},

	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarUsuario,
		RequerAutenticacao: false,
	},

	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarUsuario,
		RequerAutenticacao: false,
	},
}
