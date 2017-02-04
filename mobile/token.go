package token

import (
        "github.com/provivus/jose/crypto"
        "github.com/provivus/jose/jws"
        "github.com/provivus/jose/jwt"
)

func CreateToken() {
  // Create JWS claims
	claims := jws.Claims{}
	claims.SetAudience("example.com", "api.example.com")

	token := jws.NewJWT(claims, crypto.SigningMethodHS256)
	serializedToken, _ := token.Serialize([]byte("abcdef"))
}
