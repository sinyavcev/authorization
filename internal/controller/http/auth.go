package http

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sinyavcev/authorization/internal/models/entity/backendModels"
)

func (h *HttpController) signUp(w http.ResponseWriter, r *http.Request) {
	var (
		requestData backendModels.SignUpRequest
	)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Errorf("io.ReadAll: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &requestData); err != nil {
		h.logger.Errorf("json.Unmarshal: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := h.backendUsecases.SignUp(requestData)
	if err != nil {
		h.logger.Errorf("backendUsecases.SignUp: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	JSONResponse(w, http.StatusOK, data)
}

func (h *HttpController) signIn(w http.ResponseWriter, r *http.Request) {
	var (
		requestData backendModels.SignInRequest
	)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Errorf("io.ReadAll: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &requestData); err != nil {
		h.logger.Errorf("json.Unmarshal: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := h.backendUsecases.SignIn(requestData)
	if err != nil {
		h.logger.Errorf("backendUsecases.SignIn: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	JSONResponse(w, http.StatusOK, data)
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
