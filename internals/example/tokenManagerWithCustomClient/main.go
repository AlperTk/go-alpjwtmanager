package main

import (
	"crypto/tls"
	jwt_token_creator "github.com/AlperTk/go-alpjwtmanager/pkg/v3"
	"net/http"
)

func main() {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // Disables unknown certificate verification, it's not safe use this for testing purposes
	}
	client := &http.Client{Transport: tr}

	var tokenManager = jwt_token_creator.NewTokenManagerWithCustomClient(
		"https://localhost:8443/auth/realms/marsrealm/protocol/openid-connect/token",
		"vpncontroller",
		"BMxjKIYZxqc3rJwWEci8TPO40mjVccls",
		client,
	)

	_, _ = tokenManager.GetBearerToken()
}
