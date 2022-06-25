package main

import (
	"crypto/tls"
	"github.com/AlperTk/jwt-token-creator/src/impl/keycloak"
	"net/http"
	"time"
)

func main() {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // <--- Problem
	}
	client := &http.Client{Transport: tr}

	var tokenManager = keycloak.NewTokenManagerWithCustomClient(
		"https://localhost:8443/auth/realms/marsrealm/protocol/openid-connect/token",
		"vpncontroller",
		"BMxjKIYZxqc3rJwWEci8TPO40mjVccls",
		client,
	)
	for {
		tokenManager.GetBearerToken()
		time.Sleep(2 * time.Second)
	}
}
