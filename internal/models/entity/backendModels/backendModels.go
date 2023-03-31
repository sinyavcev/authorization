package backendModels

type UserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type UserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
