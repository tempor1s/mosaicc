package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/tempor1s/mosaic/config"
)

// Jwks represents an array of json web keys
type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

// JSONWebKeys is a struct representation of a json web key
type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

// getJwtMiddleware will return the jwt token middleware configuration
func getJwtMiddleware() jwtmiddleware.JWTMiddleware {
	c := config.GetConfig()
	return *jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			// Verify 'aud' claim
			aud := c.Auth0Domain
			checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
			if !checkAud {
				return token, errors.New("invalid audience")
			}
			// Verify 'iss' claim
			iss := c.Auth0ID
			checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
			if !checkIss {
				return token, errors.New("invalid issuer")
			}

			cert, err := getPemCert(token)
			if err != nil {
				log.Fatal(err)
			}

			result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			return result, nil
		},
		SigningMethod: jwt.SigningMethodRS256,
	})
}

// getPemCert will get the pem certificate from the API
func getPemCert(token *jwt.Token) (string, error) {
	c := config.GetConfig()
	cert := ""
	url := fmt.Sprintf("%s.well-known/jwks.json", c.Auth0ID)
	log.Println(url)
	resp, err := http.Get(url)

	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks Jwks
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	log.Println(err)
	if err != nil {
		return cert, err
	}

	for k := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("unable to find appropriate key")
		return cert, err
	}

	return cert, nil
}
