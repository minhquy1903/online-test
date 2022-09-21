package presenter

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Id       string `json:"id"`
	Name string `json:"name"`
}