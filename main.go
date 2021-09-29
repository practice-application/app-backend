package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"github.com/practice-application/app-backend/auth"
	"github.com/practice-application/app-backend/handler"
	"github.com/practice-application/app-backend/store"
)

var tokenAuth *auth.JWTAuth

func init() {
	jwks, _ := auth.JKS("https://dev-k6bx05vf.us.auth0.com/.well-known/jwks.json")
	tokenAuth = auth.New("RS256", jwks)
}

func main() {
	s := store.Connect()

	// chi
	r := chi.NewRouter()
	r.Use(
		middleware.Logger,
		middleware.StripSlashes,
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "QUERY"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
			// Debug:            true,
		}),
		auth.Verifier(tokenAuth),
		auth.Authenticator,
	)

	p := &handler.Person{
		Store: s,
	}
	r.Route("/people", func(r chi.Router) {
		r.With(auth.Authz("write:people")).Post("/", p.Create)
		r.With(auth.Authz("read:people")).Get("/{id}", p.Get)
		r.With(auth.Authz("read:people")).Get("/", p.Query)
		r.With(auth.Authz("write:people")).Put("/{id}", p.Update)
		r.With(auth.Authz("write:people")).Delete("/{id}", p.Delete)
	})

	o := &handler.Org{
		Store: s,
	}
	r.Route("/organisations", func(r chi.Router) {
		r.With(auth.Authz("write:organisations")).Post("/", o.Create)
		r.With(auth.Authz("read:organisations")).Get("/{id}", o.Get)
		r.With(auth.Authz("read:organisations")).Get("/", o.Query)
		r.With(auth.Authz("write:organisations")).Put("/{id}", o.Update)
		r.With(auth.Authz("write:organisations")).Delete("/{id}", o.Delete)
	})

	prd := &handler.Product{
		Store: s,
	}
	r.Route("/products", func(r chi.Router) {
		r.With(auth.Authz("write:products")).Post("/", prd.Create)
		r.With(auth.Authz("read:products")).Get("/{id}", prd.Get)
		r.With(auth.Authz("read:products")).Get("/", prd.Query)
		r.With(auth.Authz("write:products")).Put("/{id}", prd.Update)
		r.With(auth.Authz("write:products")).Delete("/{id}", prd.Delete)
	})

	// start server
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), r); err != nil {
		fmt.Print(err)
	}
}
