package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dansa-sociais/api/banco"
	"github.com/dansa-sociais/api/internal/autenticacao"
	"github.com/dansa-sociais/api/internal/entity"
	"github.com/dansa-sociais/api/internal/response"
	"github.com/dansa-sociais/api/internal/seguranca"
	"github.com/dansa-sociais/api/internal/usecase/repo"
)

func Login(w http.ResponseWriter, r *http.Request) {

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
	}

	var usuario entity.Usuario

	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repo.NovoRepositorioDeUsuario(db)

	usuarioSalvoNoBanco, erro := repositorio.BuscarPorEmail(usuario.Email)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro := seguranca.VerificarSenha(usuarioSalvoNoBanco.Senha, usuario.Senha); erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := autenticacao.CriarToken(usuarioSalvoNoBanco.ID)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}

	w.Write([]byte(token))
}
