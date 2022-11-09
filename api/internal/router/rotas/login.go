package rotas

import (
	"net/http"

	"github.com/dansa-sociais/api/internal/controllers"
)

var rotaLogin = Rota{
	URI:                "/login",
	Metodo:             http.MethodPost,
	Funcao:             controllers.Login,
	RequerAutenticacao: false,
}
