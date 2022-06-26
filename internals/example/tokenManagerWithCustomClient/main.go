package main

import (
	"crypto/tls"
	"github.com/AlperTk/go-alpjwtmanager/src/impl/keycloak"
	"net/http"
)

func main() {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // Disables unknown certificate verification, it's not safe use this for testing purposes
	}
	client := &http.Client{Transport: tr}

	var tokenManager = keycloak.NewTokenManagerWithCustomClient(
		"https://localhost:8443/auth/realms/marsrealm/protocol/openid-connect/token",
		"vpncontroller",
		"BMxjKIYZxqc3rJwWEci8TPO40mjVccls",
		client,
	)

	_, _ = tokenManager.GetBearerToken()
}
