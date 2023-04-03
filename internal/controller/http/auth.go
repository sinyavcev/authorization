package http

import (
	"encoding/json"
	"fmt"
	"github.com/sinyavcev/authorization/internal/models/entity/backendModels"
	"io"
	"net/http"
)

func JSONResponse(res http.ResponseWriter, message string, status int, data interface{}) {

	var Response = backendModels.Response{
		Code:     status,
		Message:  message,
		Response: data,
	}
	JSONResponse, jsonError := json.Marshal(Response)

	if jsonError != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Marshal error"))
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(Response.Code)
	res.Write(JSONResponse)

}

func (c *Controller) signUp(res http.ResponseWriter, req *http.Request) {
	var (
		user         backendModels.UserRequest
		userResponse backendModels.UserResponse
	)

	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Errorf("parse body: %w", err)
	}

	if err := json.Unmarshal(body, &user); err != nil {
		JSONResponse(res, err.Error(), http.StatusInternalServerError, nil)
	}

	userResponse, err = c.BackendUsecases.Signup(user)

	if err != nil {
		fmt.Errorf("error: %w", err)
	}

	JSONResponse(res, "success", http.StatusOK, userResponse)
}

func (c *Controller) signIn(res http.ResponseWriter, req *http.Request) {
	var (
		user         backendModels.UserRequest
		userResponse backendModels.UserResponse
	)

	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Errorf("parse body: %w", err)
	}

	if err := json.Unmarshal(body, &user); err != nil {
		JSONResponse(res, err.Error(), http.StatusInternalServerError, nil)
		return
	}

	userResponse, err := c.BackendUsecases.Signin(user)

	if err != nil {
		JSONResponse(res, err.Error(), http.StatusInternalServerError, nil)
		return
	}

	JSONResponse(res, "success", http.StatusOK, userResponse)
}

func (c *Controller) refreshToken(w http.ResponseWriter, req *http.Request) {

}

func (c *Controller) logout(w http.ResponseWriter, req *http.Request) {

}

func (c *Controller) me(w http.ResponseWriter, req *http.Request) {

}
