package repo

import (
	"database/sql"
	"fmt"

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

func (repositorio usuarios) BuscarPorNomeOuNick(nomeOuNick string) ([]entity.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) //%nomeOuNick%

	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, dataCriacao from usuarios where nome LIKE ? or nick LIKE ?",
		nomeOuNick, nomeOuNick,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []entity.Usuario

	for linhas.Next() {
		var usuario entity.Usuario
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.DataCriacao,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repositorio usuarios) BuscarUsuario(ID uint64) (entity.Usuario, error) {

	linhas, erro := repositorio.db.Query("select id, nome, nick, email, dataCriacao from usuarios where id = ?", ID)

	if erro != nil {
		return entity.Usuario{}, erro
	}

	defer linhas.Close()

	var usuario entity.Usuario

	for linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.DataCriacao,
		); erro != nil {
			return entity.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (repositorio usuarios) AtualizarUsuario(ID uint64, usuarioAtualizado entity.Usuario) error {

	statement, erro := repositorio.db.Prepare("update usuarios set nome = ?, nick = ?, email = ? where id = ?")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(usuarioAtualizado.Nome, usuarioAtualizado.Nick, usuarioAtualizado.Email, ID); erro != nil {
		return erro
	}

	return nil

}
