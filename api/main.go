package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dansa-sociais/api/rsc/config"
	"github.com/dansa-sociais/api/rsc/router"
)

func main() {

	config.Carregar()

	fmt.Sprintf("Servidor iniciado com sucesso na porta %d", config.Porta)
	r := router.Gerar()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
