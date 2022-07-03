# Go-AlpJwtManager

## Usage

```golang
package main

import "github.com/AlperTk/go-alpjwtmanager/v4"

func main() {
	var tokenManager = tokenmanager.NewTokenManager(
		"https://localhost:8443/auth/realms/marsrealm/protocol/openid-connect/token",
		"vpncontroller",
		"BMxjKIYZxqc3rJwWEci8TPO40mjVccls",
	)

	_, _ = tokenManager.GetBearerToken()
}
```
