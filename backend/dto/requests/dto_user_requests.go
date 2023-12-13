package requests

type DtoUserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DtoUserRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Status   bool   `json:"status"`
	Role     string `json:"role"`
}

type DtoUserUpdateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Status   bool   `json:"status"`
}

//Token icin response class olusturdum.

type LoginResponse struct {
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
}
