package autenticacao

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dansa-sociais/api/config"
	jwt "github.com/dgrijalva/jwt-go"
)

func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{} //map que tem as permissoes dentro do token
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix() //.unix devolve o milissegundos que passaram desde o dia 01 de jan de 1970 quando começa a era unix
	permissoes["usuarioId"] = usuarioID

	// assinando secrets
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)

	return token.SignedString(config.SecretKey) // Secret
}

func ValidarToken(r *http.Request) error {
	tokenString := extratirToken(r)

	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)

	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Token invalido")
}

/*
	O token vem com a palavra Bearer na frente, espaço e o token em sí.
	Assim, o método abaixo verifica se vem apenas 2 valores e, caso sim, retorna apenas o token que se encontra na primeira posição
	caso não, retorna string vazio e o validartoken vai saber fazer a validação
*/
func extratirToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado %v", token.Header["alg"]) // o algoritmo método que foi utilizado a assinatura, fica na propriedade alg
	}

	return config.SecretKey, nil
}
