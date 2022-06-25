package tokenManager

type TokenManager interface {
	GetBearerToken() (string, error)
	RefreshToken() error
	ClearToken()
}
