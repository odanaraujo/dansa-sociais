package autenticacao

import (
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
