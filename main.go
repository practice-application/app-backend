package main

import (
	"context"
	"fmt"
	"net/http"

	// "github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"

	"github.com/el-zacharoo/goService-shared/handler"
	"github.com/el-zacharoo/goService-shared/store"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}

func main() {
	s := store.Store{}
	s.Connect()

	// chi
	r := chi.NewRouter()
	r.Use(
		middleware.Logger,
		middleware.StripSlashes,
		jwtauth.Verifier(tokenAuth),
		jwtauth.Authenticator,
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "QUERY"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		}),
	)

	p := &handler.Person{
		Store: s,
	}
	r.Route("/people", func(r chi.Router) {
		r.With(authz("read:people")).Post("/", p.Create)
		r.Get("/{id}", p.Get)
		r.Get("/", p.Query)
		r.Put("/{id}", p.Update)
		r.Delete("/{id}", p.Delete)
	})

	o := &handler.Org{
		Store: s,
	}
	r.Route("/organisations", func(r chi.Router) {
		r.Post("/", o.Create)
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
