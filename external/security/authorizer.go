package security

import (
	"errors"
	"log"
	"time"

	"github.com/alextanhongpin/pkg/gojwt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kelseyhightower/envconfig"
)

type Claims struct {
	*gojwt.Claims
}

type Authorizer interface {
	Verify(token string) (*Claims, error)
}

type JwtConfig struct {
	Audience string `envconfig:"JWT_AUDIENCE"`
	Issuer   string `envconfig:"JWT_ISSUER"`
	Secret   string `envconfig:"JWT_SECRET"`
}

func NewAuthorizerConfig() JwtConfig {
	var cfg JwtConfig
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}

type JwtAuthorizer struct {
	gojwt.Signer
}

func (authz *JwtAuthorizer) Verify(token string) (*Claims, error) {
	claims, err := authz.Signer.Verify(token)
	if err != nil {
		return nil, err
	}
	return &Claims{
		Claims: claims,
	}, nil
}

func NewAuthorizer(cfg JwtConfig) *JwtAuthorizer {
	var (
		audience     = cfg.Audience
		issuer       = cfg.Issuer
		secret       = cfg.Secret
		expiresAfter = 7 * 24 * time.Hour // 1 Week.
	)

	validator := func(c *gojwt.Claims) error {
		if c.Issuer != issuer || c.Audience != audience {
			return errors.New("jwt: invalid claims")
		}
		return nil
	}

	return &JwtAuthorizer{
		Signer: gojwt.New(
			gojwt.Option{
				Secret:       []byte(secret),
				ExpiresAfter: expiresAfter,
				DefaultClaims: &gojwt.Claims{
					StandardClaims: jwt.StandardClaims{
						Audience: audience,
						Issuer:   issuer,
					},
				},
				Validator: validator,
			},
		),
	}
}
