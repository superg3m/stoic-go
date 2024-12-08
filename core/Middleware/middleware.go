package Middleware

import (
	"fmt"
	"github.com/superg3m/stoic-go/core/Client"
	"github.com/superg3m/stoic-go/core/Server"
)

// Logging Middleware
// CORS Middleware
// Authentication Middleware (Authorized Levels) (Tree hierarchy) ((Bitwise - 1) & flag) == 1) means has auth requirement

type StoicMiddleware func(next Server.StoicHandlerFunc) Server.StoicHandlerFunc

type AuthLevel int

const (
	USER AuthLevel = 1 << iota
	MODERATOR
	ADMIN
)

func validParams(requiredParams ...string) StoicMiddleware {
	return func(next Server.StoicHandlerFunc) Server.StoicHandlerFunc {
		return func(req *Client.StoicRequest, res Server.StoicResponse) {
			if !req.HasAll(requiredParams...) {
				res.SetError(fmt.Sprintf("Missing required parameters: %v", requiredParams))
				return
			}
			next(req, res)
		}
	}
}

func CORS() StoicMiddleware {
	return func(next Server.StoicHandlerFunc) Server.StoicHandlerFunc {
		return func(req *Client.StoicRequest, res Server.StoicResponse) {
			headers := res.Header()
			headers.Add("Access-Control-Allow-Origin", "*")
			headers.Add("Vary", "Origin")
			headers.Add("Vary", "Access-Control-Request-Method")
			headers.Add("Vary", "Access-Control-Request-Headers")
			headers.Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token")
			headers.Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

			next(req, res)
		}
	}
}

func Oauth() StoicMiddleware {
	return func(next Server.StoicHandlerFunc) Server.StoicHandlerFunc {
		return func(req *Client.StoicRequest, res Server.StoicResponse) {
			token := req.GetStringParam("oauth_token")
			if token == "" || !isValidOauthToken(token) {
				res.SetError("Invalid or missing OAuth token")
				return
			}
			next(req, res)
		}
	}
}

func isValidOauthToken(token string) bool {
	return token == "valid-oauth-token"
}

func JWT() StoicMiddleware {
	return func(next Server.StoicHandlerFunc) Server.StoicHandlerFunc {
		return func(req *Client.StoicRequest, res Server.StoicResponse) {
			token := req.GetStringParam("jwt")
			if token == "" || !isValidJWT(token) {
				res.SetError("Invalid or missing JWT token")
				return
			}
			next(req, res)
		}
	}
}

func isValidJWT(token string) bool {
	return token == "valid-jwt"
}

func logger() StoicMiddleware {
	return func(next Server.StoicHandlerFunc) Server.StoicHandlerFunc {
		return func(req *Client.StoicRequest, res Server.StoicResponse) {
			fmt.Printf("Received request: Method=%s, Path=%s\n", req.Request.Method, req.Request.URL.Path)
			next(req, res)
		}
	}
}
