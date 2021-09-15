package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"

	"github.com/practice-application/app-backend/model"
	"github.com/practice-application/app-backend/store"
)

type Person struct {
	Store *store.Store
}

func (p *Person) Create(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		return
	}

	defer r.Body.Close()
	reqByt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("err %v", err)))
	}
	var psn model.Person
	json.Unmarshal(reqByt, &psn)

	psn.ID = uuid.New().String()
	psn.Date = time.Now().String()
	p.Store.AddPerson(psn)
	w.Write([]byte("done"))
}

func (p *Person) Get(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		return
	}

	id := chi.URLParam(r, "id")

	psn, err := p.Store.GetPerson(id)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %v", err)))
	}

	rspByt, err := json.Marshal(psn)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %v", err)))
	}

	w.Write(rspByt)
}

func (p *Person) Query(w http.ResponseWriter, r *http.Request) {

	fn := r.URL.Query().Get("fn")
	ln := r.URL.Query().Get("ln")
	st := r.URL.Query().Get("st")
	lmt := int64(10)
	skip := int64(10)

	ppl, err := p.Store.GetPeople(fn, ln, st, &lmt, &skip)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %v", err)))
	}

	rspByt, err := json.Marshal(ppl)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %v", err)))
	}
	w.Write(rspByt)
}

func (p *Person) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	reqByt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("err %v", err)))
	}
	var psn model.Person
	json.Unmarshal(reqByt, &psn)

	id := chi.URLParam(r, "id")

	p.Store.UpdatePerson(id, psn)
	w.Write([]byte("done"))
}

func (p *Person) Delete(w http.ResponseWriter, r *http.Request) {

	// parts := strings.Split(r.RequestURI, "/")
	// id := parts[len(parts)-1]
	id := chi.URLParam(r, "id")

	if err := p.Store.DeletePerson(id); err != nil {
		w.Write([]byte(fmt.Sprintf("error %v", err)))
	}
	w.Write([]byte("done"))
}
