package eventService

import (
	"encoding/json"
	eventModel "g-oriekhov/testProject1/models"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"
)

func JsonResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
}

func GetEvents(w http.ResponseWriter, r *http.Request) {
	data, _ := eventModel.GetAll()
	JsonResponse(w, data)
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := eventModel.GetOne(int(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	JsonResponse(w, data)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := eventModel.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	JsonResponse(w, data)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var data eventModel.EventModel
	json.NewDecoder(r.Body).Decode(&data)
	data.Id = eventModel.GetUniqueId()
	resp, err := eventModel.Save(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	JsonResponse(w, resp)
}
