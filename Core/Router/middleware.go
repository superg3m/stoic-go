package Router

import (
	"fmt"
	"net/http"
	"reflect"
)

// Logging Middleware
// CORS Middleware
// Authentication Middleware (Authorized Levels) (Tree hierarchy) ((Bitwise - 1) & flag) == 1) means has auth requirement

// https://www.youtube.com/watch?v=ALbAYpNC6s8

type StoicMiddleware func(next StoicHandlerFunc) StoicHandlerFunc

func MiddlewareValidParams(requiredParams ...string) StoicMiddleware {
	return func(next StoicHandlerFunc) StoicHandlerFunc {
		return func(req *StoicRequest, res *StoicResponse) {
			var retParams []string
			for _, param := range requiredParams {
				if !req.Has(param) {
					retParams = append(retParams, param)
				}
			}

			if len(retParams) != 0 {
				res.AddError(fmt.Sprintf("Missing required parameters: %v", retParams))
				return
			}

			next(req, res)
		}
	}
}

func MiddlewareCORS() StoicMiddleware {
	return func(next StoicHandlerFunc) StoicHandlerFunc {
		return func(req *StoicRequest, res *StoicResponse) {
			headers := res.Header()
			headers.Set("Access-Control-Allow-Origin", "http://localhost:5173") // Set frontend origin
			headers.Set("Access-Control-Allow-Credentials", "true")             // Allow cookies
			headers.Set("Vary", "Origin, Access-Control-Request-Method, Access-Control-Request-Headers")
			headers.Set("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token")
			headers.Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

			if req.Request.Method == "OPTIONS" {
				res.WriteHeader(http.StatusOK)
				return
			}

			next(req, res)
		}
	}
}

func MiddlewareLogger() StoicMiddleware {
	return func(next StoicHandlerFunc) StoicHandlerFunc {
		return func(req *StoicRequest, res *StoicResponse) {
			fmt.Printf("Received request: Method=%s, Path=%s\n", req.Request.Method, req.Request.URL.Path)
			next(req, res)
		}
	}
}

func MiddlewareRegisterCommon(middlewares ...StoicMiddleware) {
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
