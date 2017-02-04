package token

import (
        "github.com/SermoDigital/jose/crypto"
        "github.com/SermoDigital/jose/jws"
        "github.com/SermoDigital/jose/jwt"
)

func CreateToken(s string) string {
  // Create JWS claims
	claims := jws.Claims{}
	claims.SetAudience("example.com", "api.example.com")

	token := jws.NewJWT(claims, crypto.SigningMethodES256)

	serializedToken, _ := token.Serialize()
  
  //.Serialize([]byte("abcdef"))
  return serializedToken
}
