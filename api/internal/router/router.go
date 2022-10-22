package router

import (
	"github.com/dansa-sociais/api/internal/router/rotas"
	"github.com/gorilla/mux"
)

//vai gera um router com as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
