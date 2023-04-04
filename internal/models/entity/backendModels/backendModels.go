package backendModels

type SignInResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}
type SignUpResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type SignInRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUpRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}
