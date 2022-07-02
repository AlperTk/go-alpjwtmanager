package tokenmanager

type AlpTokenManager interface {
	GetBearerToken() (string, error)
	RefreshToken() error
	ClearToken()
}
