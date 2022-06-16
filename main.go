package main

import (
	"fmt"
	"net/url"

	"github.com/google/uuid"
	verifier "github.com/nirasan/go-oauth-pkce-code-verifier"
)

const keycloakUrl = "http://localhost:8080"

func main() {
	uuidState := uuid.New()
	uuidNonce := uuid.New()
	v, _ := verifier.CreateCodeVerifier()
	urlString := fmt.Sprintf("%s/auth/realms/find-psy/protocol/openid-connect/auth"+
		"?client_id=%s"+
		"&redirect_uri=%s"+
		"&state=%s"+
		"&response_mode=fragment"+
		"&response_type=code"+
		"&scope=openid"+
		"&nonce=%s"+
		"&code_challenge=%s"+
		"&code_challenge_method=S256"+
		"&kc_idp_hint=google",
		keycloakUrl,
		"account-console",
		url.QueryEscape("http://localhost:8080/auth/realms/find-psy/account/#/personal-info"),
		uuidState.String(),
		uuidNonce.String(),
		v.CodeChallengeS256(),
	)
	fmt.Println(urlString)
}
