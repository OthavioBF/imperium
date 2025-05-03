package api

import (
	"net/http"

	"github.com/othavioBF/imperium/internal/infra/pgstore"
	"github.com/othavioBF/imperium/internal/jsonutils"
)

func (api *Api) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJson[user.GetUserByIdReq](r)
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusBadRequest, problems)
		return
	}

	user, err := api.UserService.GetUsers(r.Context())
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusBadRequest, map[string]string{
			"message": err,
		})
		return
	}

	jsonutils.EncodeJson(w, r, http.StatusOK, user)
}

func (api *Api) handleGetUserById(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJson[user.GetUserByIdReq](r)
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusBadRequest, problems)
		return
	}

	user, err := api.UserService.GetUserById(r.Context(), data.Id)
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusBadRequest, map[string]string{
			"message": err,
		})
		return
	}

	jsonutils.EncodeJson(w, r, http.StatusOK, user)
}

func (api *Api) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	_, problems, err := jsonutils.DecodeValidJson[pgstore.CreateUserParams](r)
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusBadRequest, problems)
		return
	}

	id, err := api.UserService.CreateUser(r.Context())
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusBadRequest, map[string]string{
			"message": err,
		})
		return
	}

	jsonutils.EncodeJson(w, r, http.StatusBadRequest, map[string]string{})
}

func (api *Api) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	_, problems, err := jsonutils.DecodeValidJson[pgstore.CreateUserParams](r)
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusBadRequest, problems)
		return
	}

	id, err := api.UserService.CreateUser(r.Context())
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusBadRequest, map[string]string{
			"message": err,
		})
		return
	}

	jsonutils.EncodeJson(w, r, http.StatusBadRequest, map[string]string{})
}

func (api *Api) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	_, problems, err := jsonutils.DecodeValidJson[pgstore.CreateUserParams](r)
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusBadRequest, problems)
		return
	}

	id, err := api.UserService.CreateUser(r.Context())
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusBadRequest, map[string]string{
			"message": err,
		})
		return
	}

	jsonutils.EncodeJson(w, r, http.StatusBadRequest, map[string]string{})
}
