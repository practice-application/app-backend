package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/jwk"

	"github.com/practice-application/app-backend/auth"
	"github.com/practice-application/app-backend/handler"
	"github.com/practice-application/app-backend/store"
)

var tokenAuth *auth.JWTAuth
var jwtMiddleware *jwtmiddleware.JWTMiddleware

func init() {

	keyset, _ := getJKS()
	//key, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
	tokenAuth = auth.New("RS256", keyset)

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims:
	// _, tokenString, _ := tokenAuth.Encode(map[string]interface{}{
	// 	"permissions": []interface{}{
	// 		"write:people",
	// 		"read:people",
	// 	},
	// })
	// fmt.Printf("DEBUG sample jwt: %s\n\n", tokenString)
}

func getJKS() (jwk.Set, error) {
	resp, err := http.Get("https://dev-k6bx05vf.us.auth0.com/.well-known/jwks.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return jwk.Parse(byt)
}

func Verifier(ja string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if err := jwtMiddleware.CheckJWT(w, r); err != nil {
				http.Error(w, "insufficient permission", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func main() {
	s := store.Store{}
	s.Connect()

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
			Debug:            true,
		}),
		// Verifier(""),
		auth.Verifier(tokenAuth),
		auth.Authenticator,
	)

	p := &handler.Person{
		Store: s,
	}
	r.Route("/people", func(r chi.Router) {
		r.With(authz("write:people")).Post("/", p.Create)
		r.With(authz("read:people")).Get("/{id}", p.Get)
		r.Get("/", p.Query)
		r.Put("/{id}", p.Update)
		r.Delete("/{id}", p.Delete)
	})

	o := &handler.Org{
		Store: s,
	}
	r.Route("/organisations", func(r chi.Router) {
		r.Post("/", o.Create)
		r.Get("/", o.Query)
		r.Get("/{id}", o.Get)
		r.Put("/{id}", o.Update)
		r.Delete("/{id}", o.Delete)
	})

	prd := &handler.Product{
		Store: s,
	}
	r.Route("/product", func(r chi.Router) {
		r.Post("/", prd.Create)
		r.Get("/{id}", prd.Get)
		r.Put("/{id}", prd.Update)
		r.Delete("/{id}", prd.Delete)
	})

	// start server
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Print(err)
	}
}

func authz(p string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if allow := authorise(r.Context(), p); !allow {
				http.Error(w, "insufficient permission", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func authorise(ctx context.Context, permission string) bool {
	_, claims, _ := jwtauth.FromContext(ctx)
	ps, ok := claims["permissions"].([]interface{})
	if !ok {
		return false
	}

	allow := false
	for _, p := range ps {
		if p.(string) == permission {
			allow = true
			break
		}
	}

	return allow
}
