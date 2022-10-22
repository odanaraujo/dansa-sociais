package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dansa-sociais/api/banco"
	"github.com/dansa-sociais/api/internal/entity"
	"github.com/dansa-sociais/api/internal/response"
	"github.com/dansa-sociais/api/internal/usecase/repo"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
	}

	var usuario entity.Usuario

	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
	}

	if erro = usuario.Preparar(); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}

	defer db.Close()

	repositorio := repo.NovoRepositorioDeUsuario(db)

	erro = repositorio.Save(usuario)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}

	response.JSON(w, http.StatusCreated, usuario)
}

func BuscarTodosUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuários"))
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuário"))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário"))
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário"))
}
