# Go-AlpJwtManager

## Usage

```golang
package main

import (
	"github.com/AlperTk/go-alpjwtmanager/src/impl/keycloak"
)

func main() {
	var tokenManager = keycloak.NewTokenManager(
		"https://localhost:8443/auth/realms/marsrealm/protocol/openid-connect/token",
		"vpncontroller",
		"BMxjKIYZxqc3rJwWEci8TPO40mjVccls",
	)

	bearerToken, err = tokenManager.GetBearerToken()
}
```
