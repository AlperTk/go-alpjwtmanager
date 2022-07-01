package jwt_token_creator

type TokenResponseModel struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int64  `json:"expires_in"`
	RefreshExpiresIn int64  `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	SessionState     string `json:"session_state"`
	TokenType        string `json:"token_type"`
	ExpiredAt        int64
	RefreshExpiredAt int64
}
