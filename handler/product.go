package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/el-zacharoo/goService-shared/data"
	"github.com/el-zacharoo/goService-shared/store"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type Product struct {
	Store store.Store
}

func (prd *Product) Create(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	reqByt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("err %v", err)))
	}
	var prod data.Product
	json.Unmarshal(reqByt, &prod)

	prod.ID = uuid.New().String()
	prd.Store.AddProduct(prod)
	w.Write([]byte("done"))
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
	var prod data.Product
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
