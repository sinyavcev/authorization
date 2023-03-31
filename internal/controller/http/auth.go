package http

import (
	"encoding/json"
	"github.com/sinyavcev/authorization/internal/models/entity/backendModels"
	"net/http"
)

func (c *Controller) signIn(res http.ResponseWriter, req *http.Request) {
	var user backendModels.UserRequest
	json.NewDecoder(req.Body).Decode(&user)

}

func (c *Controller) signUp(res http.ResponseWriter, req *http.Request) {
	var user backendModels.UserRequest
	json.NewDecoder(req.Body).Decode(&user)

	user, err := c.BackendUsecases.Signup(user)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"message":"` + err + `"}`))
		return
	}
	json.NewEncoder(res).Encode(user)
}

func (c *Controller) refreshToken(w http.ResponseWriter, req *http.Request) {

}

func (c *Controller) logout(w http.ResponseWriter, req *http.Request) {

}

func (c *Controller) me(w http.ResponseWriter, req *http.Request) {

}
