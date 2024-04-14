package transport

import (
	"OzonTech/internal/models"
	"OzonTech/internal/services"
	"encoding/json"
	"net/http"
)

type data struct {
	Data string `json:"data"`
}

func post(w http.ResponseWriter, r *http.Request) {
	str := data{}
	err := json.NewDecoder(r.Body).Decode(&str)
	if err != nil {
		http.Error(w, "Ошибка при декодировании запроса", http.StatusBadRequest)
		return
	}
	if services.Valid(str.Data) == true {
		rez := services.HashLink(str.Data)
		err := services.AddToMap(str.Data, rez)
		if err == nil {
			json, err := json.Marshal(rez)
			if err != nil {
				http.Error(w, "Ошибка при кодировании данных в JSON", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(json)
		}
	} else {
		http.Error(w, "invalid url", http.StatusBadRequest)
		return
	}
	d := services.DbConnection()
	defer d.Close()
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
