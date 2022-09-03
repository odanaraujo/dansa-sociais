package router

import "github.com/gorilla/mux"

//vai gera um router com as rotas configuradas
func Gerar() *mux.Router {
	return mux.NewRouter()
}
