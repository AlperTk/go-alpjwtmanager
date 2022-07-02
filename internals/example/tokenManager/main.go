package main

import jwt_token_creator "github.com/AlperTk/go-alpjwtmanager/pkg/v3"

func main() {
	var tokenManager = jwt_token_creator.NewTokenManager(
		"https://localhost:8443/auth/realms/marsrealm/protocol/openid-connect/token",
		"vpncontroller",
		"BMxjKIYZxqc3rJwWEci8TPO40mjVccls",
	)

	_, _ = tokenManager.GetBearerToken()
}
