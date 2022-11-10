package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dansa-sociais/api/config"
	"github.com/dansa-sociais/api/internal/router"
)

// func init() {
// 	chave := make([]byte, 64)
// 	if _, erro := rand.Read(chave); erro != nil {
// 		log.Fatal(erro)
// 	}

// 	//string que vamos usar para assinar o token
// 	stringBase64 := base64.StdEncoding.EncodeToString(chave)

// 	fmt.Println(stringBase64)
// }

func main() {

	config.Carregar()
	fmt.Sprintf("Servidor iniciado com sucesso na porta %d", config.Porta)

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
