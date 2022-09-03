package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dansa-sociais/api/rsc/router"
)

func main() {

	r := router.Gerar()

	fmt.Println("Servidor iniciado com sucesso na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}
