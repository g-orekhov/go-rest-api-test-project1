package eventService

import (
	"encoding/json"
	db "g-oriekhov/testProject1/models"
	"log"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"
)

func JsonResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
}

func GetEvents(w http.ResponseWriter, r *http.Request) {
	data, _ := db.GetDB().EventGetAll()
	JsonResponse(w, data)
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	data, err := db.GetDB().EventGetOne(int(id))
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusNotFound)
		return
	}
	JsonResponse(w, data)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	data, err := db.GetDB().EventDelete(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusNotFound)
		return
	}
	JsonResponse(w, data)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	event := new(db.Event)
	json.NewDecoder(r.Body).Decode(event)
	ret, err := db.GetDB().EventCreate(event)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	JsonResponse(w, ret)
}
