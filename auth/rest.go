package auth

import (
	"net/http"
	"github.com/go-zoo/bone"
)

type RestAuthDecorator struct {
	authHandler *AuthenticationHandler
}

func (a *RestAuthDecorator) DecorateHandler(orig http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		authenticated, err := a.authHandler.IsAuthenticated(req)
		if err != nil {
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(503)
			res.Write([]byte("{\"msg\": \"service unavailable\"}"))
		} else if ! authenticated {
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(403)
			res.Write([]byte("{\"msg\": \"not authenticated\"}"))
		} else {
			orig(res, req)
		}
	}
}

func (a *RestAuthDecorator) RegisterRoutes(mux *bone.Mux) error {
	return nil
}