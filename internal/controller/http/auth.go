package http

import (
	"encoding/json"
	"github.com/sinyavcev/authorization/internal/models/entity/backendModels"
	"io"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	var Response = backendModels.Response{
		Code:     status,
		Response: data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(Response.Code)

	resp, err := json.Marshal(Response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Marshal error"))
		return
	}
	w.Write(resp)
}

func (c *Controller) signUp(w http.ResponseWriter, r *http.Request) {
	var (
		requestData backendModels.UserRequest
	)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &requestData); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := c.BackendUsecases.Signup(requestData)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if data != nil {
		JSONResponse(w, http.StatusOK, data)
	}
}

func (c *Controller) signIn(w http.ResponseWriter, r *http.Request) {
	var (
		requestData backendModels.UserRequest
	)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &requestData); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := c.BackendUsecases.Signin(requestData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if data != nil {
		JSONResponse(w, http.StatusOK, data)
	}
}

func (c *Controller) refreshToken(w http.ResponseWriter, req *http.Request) {

}

func (c *Controller) logout(w http.ResponseWriter, req *http.Request) {

}

func (c *Controller) me(w http.ResponseWriter, req *http.Request) {
}
