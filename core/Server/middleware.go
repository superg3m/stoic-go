package Server

import (
	"fmt"
	"github.com/superg3m/stoic-go/core/Client"
	"net/http"
	"reflect"
)

// Logging Middleware
// CORS Middleware
// Authentication Middleware (Authorized Levels) (Tree hierarchy) ((Bitwise - 1) & flag) == 1) means has auth requirement

// https://www.youtube.com/watch?v=ALbAYpNC6s8

type StoicMiddleware func(next StoicHandlerFunc) StoicHandlerFunc

type AuthLevel int

const (
	USER AuthLevel = 1 << iota
	MODERATOR
	ADMIN
)

func MiddlewareValidParams(requiredParams ...string) StoicMiddleware {
	return func(next StoicHandlerFunc) StoicHandlerFunc {
		return func(req *Client.StoicRequest, res StoicResponse) {
			if !req.HasAll(requiredParams...) {
				res.SetError(fmt.Sprintf("Missing required parameters: %v", requiredParams))
				return
			}
			next(req, res)
		}
	}
}

func MiddlewareCORS() StoicMiddleware {
	return func(next StoicHandlerFunc) StoicHandlerFunc {
		return func(req *Client.StoicRequest, res StoicResponse) {
			// Add CORS headers
			headers := res.Header()
			headers.Add("Access-Control-Allow-Origin", "*")
			headers.Add("Vary", "Origin")
			headers.Add("Vary", "Access-Control-Request-Method")
			headers.Add("Vary", "Access-Control-Request-Headers")
			headers.Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token")
			headers.Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

			if req.Request.Method == "OPTIONS" {
				res.WriteHeader(http.StatusOK)
				return
			}

			next(req, res)
		}
	}
}

func MiddlewareOauth() StoicMiddleware {
	return func(next StoicHandlerFunc) StoicHandlerFunc {
		return func(req *Client.StoicRequest, res StoicResponse) {
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

func MiddlewareJWT() StoicMiddleware {
	return func(next StoicHandlerFunc) StoicHandlerFunc {
		return func(req *Client.StoicRequest, res StoicResponse) {
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

func MiddlewareLogger() StoicMiddleware {
	return func(next StoicHandlerFunc) StoicHandlerFunc {
		return func(req *Client.StoicRequest, res StoicResponse) {
			fmt.Printf("Received request: Method=%s, Path=%s\n", req.Request.Method, req.Request.URL.Path)
			next(req, res)
		}
	}
}

func RegisterCommonMiddleware(middlewares ...StoicMiddleware) {
	for _, middleware := range middlewares {
		if !isMiddlewareRegistered(middleware) {
			commonMiddlewares = append(commonMiddlewares, middleware)
		}
	}
}

func isMiddlewareRegistered(middleware StoicMiddleware) bool {
	for _, existingMiddleware := range commonMiddlewares {
		if reflect.DeepEqual(existingMiddleware, middleware) {
			return true
		}
	}
	return false
}

func chainMiddleware(handler StoicHandlerFunc, middlewares []StoicMiddleware) StoicHandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}
