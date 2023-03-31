package http

import (
	"encoding/json"
	"github.com/sinyavcev/authorization/internal/models/entity"
	"log"
	"net/http"
	"time"
)

func (c *Controller) signIn(res http.ResponseWriter, req *http.Request) {
	var user entity.UserRequest
	json.NewDecoder(req.Body).Decode(&user)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&dbUser)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	userPass := []byte(user.Password)
	dbPass := []byte(dbUser.Password)
	passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)
	if passErr != nil {
		log.Println(passErr)
		response.Write([]byte(`{"response":"Wrong Password!"}`))
		return
	}
	jwtToken, err := GenerateJWT()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	response.Write([]byte(`{"token":"` + jwtToken + `"}`))

}

func (c *Controller) signUp(res http.ResponseWriter, req *http.Request) {
	var user entity.UserRequest
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
