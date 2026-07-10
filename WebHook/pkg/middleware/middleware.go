package middleware

import "triggo/pkg/config"

type Middleware struct {
	Config *config.Config
}

func NewMiddleware(c *config.Config) *Middleware {
	return &Middleware{
		Config: c,
	}
}
