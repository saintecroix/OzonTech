package transport

import (
	"OzonTech/internal/models"
	"OzonTech/internal/services"
	"encoding/json"
	"net/http"
)

type data struct {
	Memory string `json:"memory"`
	Data   string `json:"data"`
}

func post(w http.ResponseWriter, r *http.Request) {
	str := data{}
	err := json.NewDecoder(r.Body).Decode(&str)
	if err != nil {
		http.Error(w, "Ошибка при декодировании запроса", http.StatusBadRequest)
		return
	}
	switch str.Memory {
	case "inmemory":
		if services.Valid(str.Data) != true {
			http.Error(w, "invalid url", http.StatusBadRequest)
			break
		}
		rez := services.HashLink(str.Data)
		err := services.AddToMap(str.Data, rez)
		if err != nil {
			http.Error(w, "ERROR: url is already exist", http.StatusBadRequest)
			break
		}
		json, err := json.Marshal(rez)
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	case "postgres":

	default:

	}
}

func get(w http.ResponseWriter, r *http.Request) {
	str := data{}
	err := json.NewDecoder(r.Body).Decode(&str)
	if err != nil {
		http.Error(w, "Ошибка при декодировании запроса", http.StatusBadRequest)
		return
	}
	s, ok := models.Inmemory[str.Data]
	if ok {
		json, err := json.Marshal(s)
		if err != nil {
			http.Error(w, "Ошибка при кодировании данных в JSON", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	} else {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
}
