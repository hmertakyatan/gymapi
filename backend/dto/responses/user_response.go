package responses

type UserLoginResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}
