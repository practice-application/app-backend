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

type Org struct {
	Store store.Store
}

func (o *Org) Create(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	reqByt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("err %v", err)))
	}
	var orgstn data.Org
	json.Unmarshal(reqByt, &orgstn)

	orgstn.ID = uuid.New().String()
	o.Store.AddOrg(orgstn)
	w.Write([]byte("done"))
}

func (o *Org) Get(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	org, err := o.Store.GetOrg(id)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %v", err)))
	}

	rspByt, err := json.Marshal(org)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %v", err)))
	}

	w.Write(rspByt)
}

func (o *Org) Update(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	reqByt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("err %v", err)))
	}
	var org data.Org
	json.Unmarshal(reqByt, &org)

	id := chi.URLParam(r, "id")

	o.Store.UpdateOrg(id, org)
	w.Write([]byte("done"))
}

func (o *Org) Delete(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	if err := o.Store.DeleteOrg(id); err != nil {
		w.Write([]byte(fmt.Sprintf("error %v", err)))
	}
	w.Write([]byte("done"))
}
