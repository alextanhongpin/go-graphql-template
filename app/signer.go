package app

import (
	"errors"
	"log"
	"time"

	"github.com/alextanhongpin/pkg/gojwt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kelseyhightower/envconfig"
)

type JwtConfig struct {
	Audience string `envconfig:"JWT_AUDIENCE"`
	Issuer   string `envconfig:"JWT_ISSUER"`
	Secret   string `envconfig:"JWT_SECRET"`
}

func NewSigner() *gojwt.JwtSigner {
	var cfg JwtConfig
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err)
	}

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
	opt := gojwt.Option{
		Secret:       []byte(secret),
		ExpiresAfter: expiresAfter,
		DefaultClaims: &gojwt.Claims{
			StandardClaims: jwt.StandardClaims{
				Audience: audience,
				Issuer:   issuer,
			},
		},
		Validator: validator,
	}
	return gojwt.New(opt)
}
