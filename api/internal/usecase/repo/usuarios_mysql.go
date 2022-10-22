package repo

import (
	"database/sql"

	"github.com/dansa-sociais/api/internal/entity"
)

//Representa um repositorio de usuario
type usuarios struct {
	db *sql.DB
}

//cria um repositório de usuário
func NovoRepositorioDeUsuario(db *sql.DB) *usuarios {
	return &usuarios{db}
}

//Criar insere um usuário no banco de dados
func (repositorio usuarios) Save(usuario entity.Usuario) error {

	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values(?, ?, ?, ?)")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	_, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)

	if erro != nil {
		return erro
	}

	return nil
}
