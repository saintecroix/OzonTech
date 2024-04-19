package transport

import (
	"OzonTech/internal/models"
	"OzonTech/internal/services"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)

type data struct {
	Memory string `json:"memory"`
	Data   string `json:"data"`
}

type result struct {
	Url string `json:"url"`
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
			return
		}
		rez := services.HashLink(str.Data)
		err := services.AddToMap(str.Data, rez)
		if err != nil {
			http.Error(w, "ERROR: url is already exist", http.StatusBadRequest)
			break
		}
		var v models.Urls
		v.Orig = str.Data
		v.Short = rez
		json, err := json.Marshal(v)
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	case "postgres":
		if services.Valid(str.Data) != true {
			http.Error(w, "invalid url", http.StatusBadRequest)
			return
		}
		db := services.DbConnection()
		defer db.Close()
		rez := services.HashLink(str.Data)
		insert, err := db.Query(fmt.Sprintf("INSERT INTO public.urls(name, shortname) VALUES ('%s', '%s');", str.Data, rez))
		if err != nil {
			http.Error(w, "incorrect data", http.StatusBadRequest)
			panic(err)
			return
		}
		defer insert.Close()
		var v models.Urls
		v.Orig = str.Data
		v.Short = rez
		json, err := json.Marshal(v)
		if err != nil {
			http.Error(w, "Ошибка при кодировании данных в JSON", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	default:
		http.Error(w, "bad memory type", http.StatusBadRequest)
		return
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	str := data{}
	err := json.NewDecoder(r.Body).Decode(&str)
	if err != nil {
		http.Error(w, "Ошибка при декодировании запроса", http.StatusBadRequest)
		return
	}
	switch str.Memory {
	case "inmemory":
		s, ok := models.Inmemory[str.Data]
		if ok {
			var v models.Urls
			v.Orig = s
			v.Short = str.Data
			json, err := json.Marshal(v)
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
	case "postgres":
		db := services.DbConnection()
		defer db.Close()
		get, err := db.Query(fmt.Sprintf("SELECT name FROM public.urls WHERE shortname = '%s';", str.Data))
		if err != nil {
			http.Error(w, "incorrect data", http.StatusBadRequest)
			return
		}
		defer get.Close()
		var v models.Urls
		v.Short = str.Data
		for get.Next() {
			err := get.Scan(&v.Orig)
			if err != nil {
				http.Error(w, "incorrect answer from database", http.StatusInternalServerError)
				return
			}
		}
		if v.Orig != "" {
			json, err := json.Marshal(v)
			if err != nil {
				http.Error(w, "Ошибка при кодировании данных в JSON", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(json)
		}
	default:
		http.Error(w, "bad memory type", http.StatusBadRequest)
		return
	}
}
