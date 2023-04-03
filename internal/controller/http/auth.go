package http

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sinyavcev/authorization/internal/models/entity/backendModels"
)

func (c *HttpController) signUp(w http.ResponseWriter, r *http.Request) {
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

	data, err := c.backendUsecases.Signin(requestData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	JSONResponse(w, http.StatusOK, data)
}

func (c *HttpController) signIn(w http.ResponseWriter, r *http.Request) {
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

	data, err := c.backendUsecases.Signin(requestData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	JSONResponse(w, http.StatusOK, data)
}

func (c *HttpController) refreshToken(w http.ResponseWriter, req *http.Request) {

}
func (c *HttpController) logout(w http.ResponseWriter, req *http.Request) {
}
func (c *HttpController) me(w http.ResponseWriter, req *http.Request) {
}

func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data != nil {
		resp, err := json.Marshal(w)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Marshal error"))
			return
		}
		w.Write(resp)
	}
}
