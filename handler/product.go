package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"

	"github.com/practice-application/app-backend/model"
	"github.com/practice-application/app-backend/store"
)

type Product struct {
	Store *store.Store
}

func (prd *Product) Create(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	reqByt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("err %v", err)))
	}
	var prod model.Product

	json.Unmarshal(reqByt, &prod)

	prod.ID = uuid.New().String()
	prod.Date = time.Now()
	prod.Image.ID = uuid.New().String()

	prd.Store.AddProduct(prod)
	w.Write([]byte("done"))
}

func (prd *Product) Query(w http.ResponseWriter, r *http.Request) {

	nm := r.URL.Query().Get("name")
	st := r.URL.Query().Get("st")
	lmtStr := r.URL.Query().Get("lmt")
	skipStr := r.URL.Query().Get("off")
	lmt, _ := strconv.ParseInt(lmtStr, 10, 64)
	skip, _ := strconv.ParseInt(skipStr, 10, 64)

	prods, err := prd.Store.GetProducts(nm, st, &lmt, &skip)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %v", err)))
	}

	rspByt, err := json.Marshal(prods)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %v", err)))
	}
	w.Write(rspByt)
}

func (prd *Product) Get(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	prod, err := prd.Store.GetProduct(id)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %v", err)))
	}

	rspByt, err := json.Marshal(prod)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %v", err)))
	}

	w.Write(rspByt)
}

func (prd *Product) Update(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	reqByt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("err %v", err)))
	}
	var prod model.Product
	json.Unmarshal(reqByt, &prod)

	id := chi.URLParam(r, "id")

	prd.Store.UpdateProduct(id, prod)
	w.Write([]byte("done"))
}

func (prd *Product) Delete(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	if err := prd.Store.DeleteProduct(id); err != nil {
		w.Write([]byte(fmt.Sprintf("error %v", err)))
	}
	w.Write([]byte("done"))
}
