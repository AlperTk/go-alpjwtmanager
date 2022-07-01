package main

import (
	"github.com/AlperTk/go-alpjwtmanager"
)

func main() {
	var tokenManager = jwt_token_creator.NewTokenManager(
		"https://localhost:8443/auth/realms/marsrealm/protocol/openid-connect/token",
		"vpncontroller",
		"BMxjKIYZxqc3rJwWEci8TPO40mjVccls",
	)

	_, _ = tokenManager.GetBearerToken()
}
