package jwt_token_creator

type AlpTokenManager interface {
	GetBearerToken() (string, error)
	RefreshToken() error
	ClearToken()
}
