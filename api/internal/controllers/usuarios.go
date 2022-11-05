package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/dansa-sociais/api/banco"
	"github.com/dansa-sociais/api/internal/entity"
	"github.com/dansa-sociais/api/internal/response"
	"github.com/dansa-sociais/api/internal/usecase/repo"
	"github.com/gorilla/mux"
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

	if erro = usuario.Preparar("cadastro"); erro != nil {
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
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, erro := banco.Conectar()

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repo.NovoRepositorioDeUsuario(db)

	usuarios, erro := repositorio.BuscarPorNomeOuNick(nomeOuNick)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if len(usuarios) == 0 {
		var respo response.Resposta
		respo.Message = "Usuário não existe"
		response.JSON(w, http.StatusOK, respo)
		return
	}

	response.JSON(w, http.StatusOK, usuarios)
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioId, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)

	if erro != nil {
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

	usuario, erro := repositorio.BuscarUsuario(usuarioId)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if usuario.ID == 0 {
		var respo response.Resposta
		respo.Message = "Usuário não existe"
		response.JSON(w, http.StatusOK, respo)
		return
	}

	response.JSON(w, http.StatusOK, usuario)
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	usuarioId, erro := strconv.ParseUint(parametro["usuarioId"], 10, 64)

	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario entity.Usuario

	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("edicao"); erro != nil {
		fmt.Println("caiu aqui 149")
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}

	defer db.Close()

	repositorio := repo.NovoRepositorioDeUsuario(db)

	if erro := repositorio.AtualizarUsuario(usuarioId, usuario); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}

	response.JSON(w, http.StatusNoContent, nil)

}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário"))
}
