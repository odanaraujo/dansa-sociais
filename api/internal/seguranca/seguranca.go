package seguranca

import "golang.org/x/crypto/bcrypt"

//Hash recebe uma string e retorna uma senha hash
//DefaultCost é  o custo para gerar a senha.
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

//Compara uma senha com  uma hash e verifica se são iguais
func VerificarSenha(senha string, senhaComHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(senha), []byte(senhaComHash))
}
